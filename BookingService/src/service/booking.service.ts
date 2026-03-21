import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizeIdempotencyKey,
  getIdempotencyKeyWithLock,
} from "../repositories/booking.repository";
import {
  BadRequestError,
  InternalServerError,
  NotFoundError,
} from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";
import { createBookingDto } from "../dto/booking.dto";
import prismaClient from "../prisma/client";
import { redlock } from "../config/redis.config";
import { serverConfig } from "../config";

export async function createBookingService(bookingData: createBookingDto) {
  // Create a unique resource identifier for the booking based on hotel and user information
  const bookingResource = `Hotel:${bookingData.hotelId}`;

  const ttl = serverConfig.LOCK_TTL; // 5 seconds default

  try {
    // Acquire a distributed lock for the booking resource
    await redlock.acquire([bookingResource], ttl);

    // If lock is acquired, proceed to create the booking
    const booking = await createBooking({
      userId: bookingData.userId,
      hotelId: bookingData.hotelId,
      totalGuests: bookingData.totalGuests,
      bookingAmount: bookingData.bookingAmount,
    });

    const idempotencyKey = generateIdempotencyKey();
    await createIdempotencyKey(idempotencyKey, booking.id);
    return {
      bookingId: booking.id,
      idempotencyKey: idempotencyKey,
    };
  } catch (error) {
    console.error("Error occurred while creating booking:", error);
    throw new InternalServerError("Failed to acquire lock for booking");
  }
}

export async function confirmBookingService(idempotencyKey: string) {
  return await prismaClient.$transaction(async (tx) => {
    const idempotencyKeyData = await getIdempotencyKeyWithLock(
      tx,
      idempotencyKey,
    );

    if (!idempotencyKeyData || !idempotencyKeyData.bookingId) {
      throw new NotFoundError(
        "Booking not found for the given idempotency key",
      );
    }

    if (idempotencyKeyData.finalized) {
      throw new BadRequestError("Booking already finalized");
    }

    const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);

    await finalizeIdempotencyKey(tx, idempotencyKey);

    return booking;
  });
}

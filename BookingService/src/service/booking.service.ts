import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizeIdempotencyKey,
  getIdempotencyKey,
} from "../repositories/booking.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";

import { createBookingDto } from "../dto/booking.dto";

export async function createBookingService(bookingData: createBookingDto){
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
}

export async function confirmBookingService(idempotencyKey: string) {
  const idempotencyKeyData = await getIdempotencyKey(idempotencyKey);

  if (!idempotencyKeyData) {
    throw new NotFoundError("Invalid idempotency key");
  }

  if (idempotencyKeyData.finalized) {
    throw new BadRequestError("Booking already finalized");
  }

  const booking = await confirmBooking(idempotencyKeyData.bookingId!);

  await finalizeIdempotencyKey(idempotencyKey);

  return booking;
}

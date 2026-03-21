import { Request, Response } from "express";
import { StatusCodes } from "http-status-codes";
import {
  createBookingService,
  confirmBookingService,
} from "../service/booking.service";

export const createBookingHandler = async (req: Request, res: Response) => {
  const booking = await createBookingService(req.body);
  res.status(StatusCodes.CREATED).json({
    bookingId: booking.bookingId,
    idempotencyKey: booking.idempotencyKey,
  });
};

export const confirmBookingHandler = async (req: Request, res: Response) => {
  const booking = await confirmBookingService(
    req.params.idempotencyKey as string,
  );
  res.status(StatusCodes.OK).json(booking);
};

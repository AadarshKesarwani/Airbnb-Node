import { z } from "zod";

export const createBookingSchema = z.object({
    userId: z.number({message : "used id must be present"}),
    hotelId: z.number({message : "hotel id must be present"}),
    totalGuests: z.number({message : "total guests must be present"}).min(1, {message : "at least one guest must be present"}),
    bookingAmount: z.number({message : "booking amount must be present"}).positive(),
})
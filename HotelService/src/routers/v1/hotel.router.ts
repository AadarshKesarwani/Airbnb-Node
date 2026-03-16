import express from 'express';
import { createHotelHandler, getHotelByIdHandler } from '../../controllers/hotel.controller';
import { validateRequestBody } from '../../validators';
import { HotelSchema } from '../../validators/hotel.validator';




const hotelRouter = express.Router();

hotelRouter.post('/',
    validateRequestBody(HotelSchema),
    createHotelHandler);

hotelRouter.get('/:id', getHotelByIdHandler);

export default hotelRouter;
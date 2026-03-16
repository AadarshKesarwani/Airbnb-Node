import { NextFunction, Request, Response } from "express";
import { createHotelService, getHotelByIdService } from "../service/hotel.service";


export async function createHotelHandler(req: Request, res:Response, next: NextFunction) {
    try {
        const hotelData = req.body;
        const hotelResponse = await createHotelService(hotelData);
        res.status(201).json({
            message : "Hotel created successfully",
            data : hotelResponse,
            success : true
        });
    } catch (error) {
        next(error);
    }

}

export async function getHotelByIdHandler(req: Request, res:Response, next: NextFunction) {
    try {
        const hotelId = Number(req.params.id);
        const hotelResponse = await getHotelByIdService(hotelId);
        res.status(200).json({
            message : "Hotel retrieved successfully",
            data : hotelResponse,
            success : true
        });
    } catch (error) {
        next(error);
    }
}
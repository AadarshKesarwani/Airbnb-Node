import { NextFunction, Request, Response } from "express";
import { createHotelService, deleteHotelService, getAllHotelsService, getHotelByIdService } from "../service/hotel.service";
import { StatusCodes } from "http-status-codes";


export async function createHotelHandler(req: Request, res:Response, next: NextFunction) {
    try {
        const hotelData = req.body;
        const hotelResponse = await createHotelService(hotelData);
        res.status(StatusCodes.CREATED).json({
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
        res.status(StatusCodes.OK).json({
            message : "Hotel retrieved successfully",
            data : hotelResponse,
            success : true
        });
    } catch (error) {
        next(error);
    }
}


export async function getAllHotelsHandler(req: Request, res:Response, next: NextFunction) {
    try {
        const hotelsResponse = await getAllHotelsService();
        res.status(StatusCodes.OK).json({
            message : "Hotels retrieved successfully",
            data : hotelsResponse,
            success : true
        });
    } catch (error) {
        next(error);
    }
}

export async function deleteHotelHandler(req: Request, res:Response, next: NextFunction) {
    try {
        const hotelId = Number(req.params.id);
        const hotelResponse = await deleteHotelService(hotelId);
        res.status(StatusCodes.OK).json({
            message : "Hotel deleted successfully",
            data : hotelResponse,
            success : true
        });
    } catch (error) {
        next(error);
    }
}

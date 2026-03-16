import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import {CreateHotelDTO} from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelData: CreateHotelDTO){
    const hotel = await Hotel.create({
        name: hotelData.name,
        address: hotelData.address,
        location: hotelData.location,
        rating: hotelData.rating,
        ratingCount: hotelData.ratingCount
    });
    logger.info(`Hotel created with ID: ${hotel.id}`);
    return hotel;
}



export async function getHotelById(id: number) {
    const hotel = await Hotel.findByPk(id);

    if (!hotel) {
        logger.error(`Hotel with ID ${id} not found`);
        throw new NotFoundError(`Hotel with ID ${id} not found`);
    }

    logger.info(`Hotel found with ID: ${hotel.id}`);
    return hotel;
}


export async function getAllHotels() {
    const hotels = await Hotel.findAll();

    if(!hotels || hotels.length === 0) {
        logger.error(`No hotels found`);
        throw new NotFoundError(`No hotels found`);
    }

    logger.info(`Found ${hotels.length} hotels`);
    return hotels;
}
package hotel

import "context"

func (h HotelUsecase) DeleteHotel(ctx context.Context, hotelId int) error {
	err := h.hotelRepo.Delete(ctx, hotelId)

	if err != nil {
		h.logger.Error(err.Error())
		return err
	}
	
	return nil
}

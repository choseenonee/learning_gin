package hotel

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

func (h HotelUsecase) GetHotel(ctx context.Context, hotelId int) (model.Hotel, error) {
	ctx, cancel := context.WithTimeout(ctx, h.contextTimeout)
	defer cancel()

	hotel, err := h.hotelRepo.GetById(ctx, hotelId)
	if err != nil {
		h.logger.Error(err.Error())
		return model.Hotel{}, err
	}
	return hotel, nil
}

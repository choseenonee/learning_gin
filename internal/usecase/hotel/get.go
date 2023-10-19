package hotel

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

func (h HotelUsecase) GetHotel(ctx context.Context, hotelId int) (model.HotelWithContact, error) {
	ctx, cancel := context.WithTimeout(ctx, h.contextTimeout)
	defer cancel()

	hotelWithContact, err := h.hotelRepo.GetById(ctx, hotelId)
	if err != nil {
		h.logger.Error(err.Error())
		return model.HotelWithContact{}, err
	}

	return hotelWithContact, nil
}

package hotel

import (
	"context"
	"github.com/niumandzi/nto2022/model"
)

func (h HotelUsecase) GetAllHotels(ctx context.Context) ([]model.HotelWithContact, error) {
	ctx, cancel := context.WithTimeout(ctx, h.contextTimeout)
	defer cancel()

	hotelWithContact, err := h.hotelRepo.GetAll(ctx)
	if err != nil {
		h.logger.Error(err.Error())
		return []model.HotelWithContact{}, err
	}

	return hotelWithContact, nil
}

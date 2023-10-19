package hotel

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/niumandzi/nto2022/model"
	"regexp"
)

func (h HotelUsecase) CreateHotel(ctx context.Context, hotel model.Hotel) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, h.contextTimeout)

	defer cancel()

	err := validation.ValidateStruct(&hotel,
		validation.Field(&hotel.Name, validation.Required, validation.Length(0, 200)),
		validation.Field(&hotel.LocationId, validation.Required),
		validation.Field(&hotel.Number, validation.Required, validation.Match(regexp.MustCompile(`^\+\d{1,2}\s\(\d{3}\)\s\d{3}-\d{2}-\d{2}$`))),
		validation.Field(&hotel.WorkerId, validation.Required),
		validation.Field(&hotel.Description, validation.Required, validation.Length(0, 500)),
	)
	if err != nil {
		h.logger.Error(err.Error())
		return 0, err
	}

	id, err := h.hotelRepo.Create(ctx, hotel)
	if err != nil {
		h.logger.Error(err.Error())
		return 0, err
	}
	return id, nil
}

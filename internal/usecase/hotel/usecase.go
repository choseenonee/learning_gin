package hotel

import (
	"github.com/niumandzi/nto2022/internal/repository"
	"github.com/niumandzi/nto2022/pkg/logging"
	"time"
)

type HotelUsecase struct {
	hotelRepo      repository.HotelRepository
	contextTimeout time.Duration
	logger         logging.Logger
}

func NewHotelUsecase(hotel repository.HotelRepository, timeout time.Duration, logger logging.Logger) HotelUsecase {
	return HotelUsecase{
		hotelRepo:      hotel,
		contextTimeout: timeout,
		logger:         logger,
	}
}

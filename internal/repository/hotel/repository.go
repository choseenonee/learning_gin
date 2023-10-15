package hotel

import (
	"database/sql"
	"github.com/niumandzi/nto2022/pkg/logging"
)

type HotelRepository struct {
	db     *sql.DB
	logger logging.Logger
}

func NewHotelRepository(db *sql.DB, logger logging.Logger) HotelRepository {
	return HotelRepository{
		db:     db,
		logger: logger,
	}
}

package hotel

import (
	"context"
	"database/sql"
	"github.com/niumandzi/nto2022/model"
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

func (h HotelRepository) Create(ctx context.Context, hotel model.Hotel) (int, error) {
	result, err := h.db.ExecContext(
		ctx,
		"INSERT INTO hotel (name, location_id, number, worker_id, description) VALUES (?, ?, ?, ?, ?)",
		hotel.Name,
		hotel.LocationId,
		hotel.Number,
		hotel.WorkerId,
		hotel.Description,
	)
	if err != nil {
		h.logger.Error(err.Error())
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		h.logger.Error(err.Error())
		return 0, err
	}

	return int(id), nil
}

func (h HotelRepository) GetById(ctx context.Context, hotelId int) (model.Hotel, error) {
	row := h.db.QueryRowContext(ctx, "SELECT * FROM hotel WHERE id = ?", hotelId)

	var hotel model.Hotel

	err := row.Scan(
		&hotel.Id,
		&hotel.Name,
		&hotel.LocationId,
		&hotel.Number,
		&hotel.WorkerId,
		&hotel.Description,
	)

	if err != nil {
		h.logger.Error(err.Error())
		return model.Hotel{}, err
	}

	return hotel, nil
}

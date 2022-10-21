package booking_repo

import (
	"context"
	"gin-tuns_go_flight/database"
	booking_model "gin-tuns_go_flight/grpc/booking-grpc/model"

	"gorm.io/gorm"
)

//Embeded struct

type BookingRepository interface {
	CreateBooking(ctx context.Context, model *booking_model.Booking) (*booking_model.Booking, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (BookingRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&booking_model.Booking{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) CreateBooking(ctx context.Context, model *booking_model.Booking) (*booking_model.Booking, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

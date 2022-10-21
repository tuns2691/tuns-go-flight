package flight_repo

import (
	"context"
	"gin-tuns_go_flight/database"
	flight_model "gin-tuns_go_flight/grpc/flight-grpc/model"

	"gorm.io/gorm"
)

//Embeded struct

type FlightRepository interface {
	CreateFlight(ctx context.Context, model *flight_model.Flight) (*flight_model.Flight, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (FlightRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&flight_model.Flight{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) CreateFlight(ctx context.Context, model *flight_model.Flight) (*flight_model.Flight, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

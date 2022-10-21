package customer_repo

import (
	"context"
	"gin-tuns_go_flight/database"
	customer_model "gin-tuns_go_flight/grpc/customer-grpc/model"

	"gorm.io/gorm"
)

//Embeded struct

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, model *customer_model.Customer) (*customer_model.Customer, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&customer_model.Customer{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) CreateCustomer(ctx context.Context, model *customer_model.Customer) (*customer_model.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

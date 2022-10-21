package flight_model

import (
	"time"

	"gin-tuns_go_flight/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Flight struct {
	Id            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string    `gorm:"column:name"`
	From          string    `gorm:"column:from"`
	To            string    `gorm:"column:to"`
	DepartDate    time.Time `gorm:"column:depart_date"`
	Status        string    `gorm:"column:status"`
	AvailableSlot int32     `gorm:"column:available_slot"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (in *Flight) ToResponse() *pb.Flight {
	res := &pb.Flight{
		Id:            in.Id.String(),
		Name:          in.Name,
		From:          in.From,
		To:            in.To,
		DepartDate:    timestamppb.New(in.DepartDate),
		Status:        in.Status,
		AvailableSlot: in.AvailableSlot,
		CreatedAt:     timestamppb.New(in.CreatedAt),
		UpdatedAt:     timestamppb.New(in.UpdatedAt),
	}

	return res
}

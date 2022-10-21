package booking_model

import (
	"time"

	"gin-tuns_go_flight/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Booking struct {
	Id         uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerId string    `gorm:"column:customer_Id"`
	FlightId   string    `gorm:"column:flight_id"`
	Code       string    `gorm:"column:code"`
	BookedDate time.Time `gorm:"column:booked_date"`
	Status     string    `gorm:"column:status"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (in *Booking) ToResponse() *pb.Booking {
	res := &pb.Booking{
		Id:         in.Id.String(),
		CustomerId: in.CustomerId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		BookedDate: timestamppb.New(in.BookedDate),
		Status:     in.Status,
		CreatedAt:  timestamppb.New(in.CreatedAt),
		UpdatedAt:  timestamppb.New(in.UpdatedAt),
	}

	return res
}

package booking_handler

import (
	"context"
	booking_model "gin-tuns_go_flight/grpc/booking-grpc/model"
	booking_repo "gin-tuns_go_flight/grpc/booking-grpc/repo"
	"gin-tuns_go_flight/pb"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookingHandler struct {
	pb.UnimplementedRPCBookingServer
	bookingRepository booking_repo.BookingRepository
	mu                *sync.Mutex
}

func NewBookingHandler(bookingRepository booking_repo.BookingRepository) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepository: bookingRepository,
		mu:                &sync.Mutex{},
	}, nil
}

func (h *BookingHandler) CreateBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	req := &booking_model.Booking{
		Id:         uuid.New(),
		CustomerId: in.CustomerId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		BookedDate: time.Now(),
		Status:     in.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	out, err := h.bookingRepository.CreateBooking(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponse(), nil
}

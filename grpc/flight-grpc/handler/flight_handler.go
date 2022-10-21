package flight_handler

import (
	"context"
	flight_model "gin-tuns_go_flight/grpc/flight-grpc/model"
	flight_repo "gin-tuns_go_flight/grpc/flight-grpc/repo"
	"gin-tuns_go_flight/pb"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightHandler struct {
	pb.UnimplementedRPCFlightServer
	flightRepository flight_repo.FlightRepository
	mu               *sync.Mutex
}

func NewFlightHandler(flightRepository flight_repo.FlightRepository) (*FlightHandler, error) {
	return &FlightHandler{
		flightRepository: flightRepository,
		mu:               &sync.Mutex{},
	}, nil
}

func (h *FlightHandler) CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	req := &flight_model.Flight{
		Id:            uuid.New(),
		Name:          in.Name,
		From:          in.From,
		To:            in.To,
		DepartDate:    in.DepartDate.AsTime(),
		Status:        in.Status,
		AvailableSlot: in.AvailableSlot,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	flight, err := h.flightRepository.CreateFlight(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return flight.ToResponse(), nil
}

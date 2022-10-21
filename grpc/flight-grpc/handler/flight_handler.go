package flight_handler

import (
	"context"
	"database/sql"
	flight_model "gin-tuns_go_flight/grpc/flight-grpc/model"
	flight_repo "gin-tuns_go_flight/grpc/flight-grpc/repo"
	flight_request "gin-tuns_go_flight/grpc/flight-grpc/request"
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

func (h *FlightHandler) SearchFlight(ctx context.Context, in *pb.SearchFlightRequest) (*pb.SearchFlightResponse, error) {
	flights, err := h.flightRepository.SearchFlight(ctx, &flight_request.SearchFlightRequest{
		Id:       in.Id,
		Name:     in.Name,
		From:     in.From,
		To:       in.To,
		FromDate: in.FromDate.AsTime(),
		ToDate:   in.ToDate.AsTime(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.SearchFlightResponse{
		Flight: []*pb.Flight{},
	}

	for _, flight := range flights {
		pRes.Flight = append(pRes.Flight, flight.ToResponse())
	}

	if err != nil {
		return nil, err
	}

	return pRes, nil
}

package handler

import (
	"context"
	"gin-tuns_go_flight/grpc/customer-grpc/model"
	"gin-tuns_go_flight/grpc/customer-grpc/repo"
	"gin-tuns_go_flight/pb"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerHandler struct {
	pb.UnimplementedRPCCustomerServer
	customerRepository repo.CustomerRepository
	mu                 *sync.Mutex
}

func NewCustomerHandler(customerRepository repo.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepository: customerRepository,
		mu:                 &sync.Mutex{},
	}, nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	req := &model.Customer{
		Id:             uuid.New(),
		Role:           in.Role,
		Name:           in.Name,
		Email:          in.Email,
		PhoneNumber:    in.PhoneNumber,
		DateOfBith:     in.DateOfBith,
		IdentityCard:   in.IdentityCard,
		Address:        in.Address,
		MembershipCard: in.MembershipCard,
		Password:       in.Password,
		Status:         in.Status,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	customer, err := h.customerRepository.CreateCustomer(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return customer.ToResponse(), nil
}

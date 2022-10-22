package customer_handler

import (
	"context"
	"database/sql"
	customer_model "gin-tuns_go_flight/grpc/customer-grpc/model"
	customer_repo "gin-tuns_go_flight/grpc/customer-grpc/repo"
	customer_request "gin-tuns_go_flight/grpc/customer-grpc/request"
	"gin-tuns_go_flight/pb"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerHandler struct {
	pb.UnimplementedRPCCustomerServer
	customerRepository customer_repo.CustomerRepository
	mu                 *sync.Mutex
}

func NewCustomerHandler(customerRepository customer_repo.CustomerRepository) (*CustomerHandler, error) {
	return &CustomerHandler{
		customerRepository: customerRepository,
		mu:                 &sync.Mutex{},
	}, nil
}

func (h *CustomerHandler) FindById(ctx context.Context, in *pb.CustomerParamId) (*pb.Customer, error) {
	out, err := h.customerRepository.FindById(ctx, uuid.MustParse(in.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponse(), nil
}

func (h *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	req := &customer_model.Customer{
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

func (h *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	req, err := h.customerRepository.FindById(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Role >= 0 {
		req.Role = in.Role
	}

	if in.Name != "" {
		req.Name = in.Name
	}

	if in.Email != "" {
		req.Email = in.Email
	}

	if in.PhoneNumber != "" {
		req.PhoneNumber = in.PhoneNumber
	}

	if in.DateOfBith != "" {
		req.DateOfBith = in.DateOfBith
	}

	if in.IdentityCard != "" {
		req.IdentityCard = in.IdentityCard
	}

	if in.Address != "" {
		req.Address = in.Address
	}

	if in.MembershipCard != "" {
		req.MembershipCard = in.MembershipCard
	}

	if in.Password != "" {
		req.Password = in.Password
	}

	if in.Status >= 0 {
		req.Status = in.Status
	}

	req.UpdatedAt = time.Now()

	out, err := h.customerRepository.UpdateCustomer(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponse(), nil
}

func (h *CustomerHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	req, err := h.customerRepository.FindById(ctx, uuid.MustParse(in.CustomerId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.NewPassword != "" {
		req.Password = in.NewPassword
	}

	req.UpdatedAt = time.Now()

	custOut, err := h.customerRepository.UpdateCustomer(ctx, req)

	if custOut != nil && err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &pb.ChangePasswordResponse{
		Code:    0,
		Message: "Success",
	}

	return out, nil
}

func (h *CustomerHandler) SearchCustomer(ctx context.Context, in *pb.SearchCustomerRequest) (*pb.SearchCustomerResponse, error) {
	customers, err := h.customerRepository.SearchCustomer(ctx, &customer_request.SearchCustomerRequest{
		Name:         in.Name,
		Email:        in.Email,
		PhoneNumber:  in.PhoneNumber,
		IdentityCard: in.IdentityCard,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.SearchCustomerResponse{
		Customer: []*pb.Customer{},
	}

	for _, customer := range customers {
		pRes.Customer = append(pRes.Customer, customer.ToResponse())
	}

	if err != nil {
		return nil, err
	}

	return pRes, nil
}

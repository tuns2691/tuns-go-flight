package handler

import (
	"gin-tuns_go_flight/api/customer-api/request"
	"gin-tuns_go_flight/api/customer-api/response"
	"gin-tuns_go_flight/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.RPCCustomerClient
}

func NewCustomerHandler(customerClient pb.RPCCustomerClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	req := request.CreateCustomerRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Customer{
		Role:           req.Role,
		Name:           req.Name,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		DateOfBith:     req.DateOfBith,
		IdentityCard:   req.IdentityCard,
		Address:        req.Address,
		MembershipCard: req.MembershipCard,
		Password:       req.Password,
		Status:         req.Status,
	}

	pRes, err := h.customerClient.CreateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &response.CustomerResponse{
		Id:             pRes.Id,
		Role:           pRes.Role,
		Name:           pRes.Name,
		Email:          pRes.Email,
		PhoneNumber:    pRes.PhoneNumber,
		DateOfBith:     pRes.DateOfBith,
		IdentityCard:   pRes.IdentityCard,
		Address:        pRes.Address,
		MembershipCard: pRes.MembershipCard,
		Password:       pRes.Password,
		Status:         pRes.Status,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

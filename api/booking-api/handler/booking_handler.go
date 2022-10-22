package booking_handler

import (
	booking_request "gin-tuns_go_flight/api/booking-api/request"
	"gin-tuns_go_flight/pb"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookingHandler interface {
	CustomerBooking(c *gin.Context)
	GuestBooking(c *gin.Context)
	CancelBooking(c *gin.Context)
	BookingHistory(c *gin.Context)
	SearchBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient  pb.RPCBookingClient
	customerClient pb.RPCCustomerClient
	flightClient   pb.RPCFlightClient
}

func NewBookingHandler(
	bookingClient pb.RPCBookingClient,
	customerClient pb.RPCCustomerClient,
	flightClient pb.RPCFlightClient) BookingHandler {
	return &bookingHandler{
		bookingClient:  bookingClient,
		customerClient: customerClient,
		flightClient:   flightClient,
	}
}

func (h *bookingHandler) CustomerBooking(c *gin.Context) {
	req := booking_request.CustomerBookingRequest{}

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

	if req.Slot <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "99",
			"error":  "Booking Slot is invalid",
		})
	}

	// Gen Booking Code
	bookingCode := "VN_" + generateCode(6)

	pReq := &pb.Booking{
		CustomerId: req.CustomerId,
		FlightId:   req.FlightId,
		BookedSlot: req.Slot,
		Code:       bookingCode,
		Status:     "Active",
	}

	pRes, err := h.bookingClient.CreateBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	// Find by ID
	pReqFind := &pb.BookingParamId{
		Id: pRes.Id,
	}

	pResFind, err := h.bookingClient.FindById(c.Request.Context(), pReqFind)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pResFind,
	})
}

// --> Nguoi chua dang ky dat ve
func (h *bookingHandler) GuestBooking(c *gin.Context) {
	req := booking_request.GuestBookingRequest{}

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

	if req.Slot <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "99",
			"error":  "Booking Slot is invalid",
		})
	}

	// Kiem tra xem thong tin nguoi dung da dang ky chua ?
	pReqCus := &pb.SearchCustomerRequest{
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		IdentityCard: req.IdentityCard,
	}

	pResCus, err := h.customerClient.SearchCustomer(c.Request.Context(), pReqCus)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	pCustomerId := ""
	if pResCus != nil && len(pResCus.Customer) > 0 {
		// set customer_id
		pCustomerId = pResCus.Customer[0].Id
	} else {
		// If not existed => create new customer
		pReqCreateCust := &pb.Customer{
			Role:           0,
			Name:           req.Name,
			Email:          req.Email,
			PhoneNumber:    req.PhoneNumber,
			DateOfBith:     req.DateOfBith,
			IdentityCard:   req.IdentityCard,
			Address:        req.Address,
			MembershipCard: req.MembershipCard,
			Status:         1,
		}

		pResCreateCust, err := h.customerClient.CreateCustomer(c.Request.Context(), pReqCreateCust)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusText(http.StatusInternalServerError),
				"error":  err.Error(),
			})
			return
		}

		// set customer_id
		pCustomerId = pResCreateCust.Id
	}

	// Gen Booking Code
	bookingCode := "VN_" + generateCode(6)

	pReq := &pb.Booking{
		CustomerId: pCustomerId,
		FlightId:   req.FlightId,
		BookedSlot: req.Slot,
		Code:       bookingCode,
		Status:     "Active",
	}

	pRes, err := h.bookingClient.CreateBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	// Booking success -> minus slot START
	pReqFlight := &pb.FlightParamId{
		Id: req.FlightId,
	}

	pResFlight, err := h.flightClient.FindById(c.Request.Context(), pReqFlight)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	pResFlight.AvailableSlot = pResFlight.AvailableSlot - req.Slot

	pResFlight2, err := h.flightClient.UpdateFlight(c.Request.Context(), pResFlight)

	if err != nil && pResFlight2 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}
	// Booking success -> minus slot END

	// Find by ID
	pReqFind := &pb.BookingParamId{
		Id: pRes.Id,
	}

	pResFind, err := h.bookingClient.FindById(c.Request.Context(), pReqFind)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pResFind,
	})
}

func (h *bookingHandler) CancelBooking(c *gin.Context) {
	req := booking_request.CancelBookingRequest{}

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

	// Find by ID
	pReqFind := &pb.BookingParamId{
		Id: req.Id,
	}

	pResFind, err := h.bookingClient.FindById(c.Request.Context(), pReqFind)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	if pResFind.Id == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "99",
			"error":  "ID not existed",
		})
		return
	}

	pResFind.Status = "Cancel"

	pRes, err := h.bookingClient.UpdateBooking(c.Request.Context(), pResFind)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	// Find by ID
	pReqFindBk := &pb.BookingParamId{
		Id: pRes.Id,
	}

	pResFindBk, err := h.bookingClient.FindById(c.Request.Context(), pReqFindBk)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pResFindBk,
	})
}

func (h *bookingHandler) BookingHistory(c *gin.Context) {
	req := booking_request.ViewBookingRequest{}

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

	pReq := &pb.SearchBookingRequest{
		CustomerId: req.CustomerId,
	}

	pRes, err := h.bookingClient.SearchBooking(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pRes.Booking,
	})
}

func (h *bookingHandler) SearchBooking(c *gin.Context) {
	req := booking_request.SearchBookingRequest{}

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

	// pFromDate, _ := time.Parse("2006/01/02 15:04:05", req.FromDate)
	// pToDate, _ := time.Parse("2006/01/02 15:04:05", req.ToDate)

	pReq := &pb.SearchBookingRequest{
		CustomerId: req.CustomerId,
		FlightId:   req.FlightId,
		Code:       req.Code,
		Status:     req.Status,
		// FromDate:   timestamppb.New(pFromDate),
		// ToDate:     timestamppb.New(pToDate),
	}

	pRes, err := h.bookingClient.SearchBooking(c.Request.Context(), pReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pRes.Booking,
	})
}

func generateCode(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

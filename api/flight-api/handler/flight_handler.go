package flight_handler

import (
	flight_request "gin-tuns_go_flight/api/flight-api/request"
	flight_response "gin-tuns_go_flight/api/flight-api/response"
	"gin-tuns_go_flight/pb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.RPCFlightClient
}

func NewFlightHandler(flightClient pb.RPCFlightClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := flight_request.CreateFlightRequest{}

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

	DepartDateTime, _ := time.Parse("2006/01/02 15:04:05", req.DepartDate+" "+req.DepartTime)

	pReq := &pb.Flight{
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		DepartDate:    timestamppb.New(DepartDateTime),
		Status:        req.Status,
		AvailableSlot: req.AvailableSlot,
	}

	pRes, err := h.flightClient.CreateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &flight_response.CreateFlightResponse{
		Id:            pRes.Id,
		Name:          pRes.Name,
		From:          pRes.From,
		To:            pRes.To,
		DepartDate:    pRes.DepartDate.AsTime().Format("2006/01/02"),
		DepartTime:    pRes.DepartDate.AsTime().Format("15:04:05"),
		Status:        pRes.Status,
		AvailableSlot: pRes.AvailableSlot,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

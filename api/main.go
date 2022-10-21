package main

import (
	booking_handler "gin-tuns_go_flight/api/booking-api/handler"
	customer_handler "gin-tuns_go_flight/api/customer-api/handler"
	flight_handler "gin-tuns_go_flight/api/flight-api/handler"
	"gin-tuns_go_flight/middleware"
	"gin-tuns_go_flight/pb"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	conn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	customerClient := pb.NewRPCCustomerClient(conn)
	bookingClient := pb.NewRPCBookingClient(conn)
	flightClient := pb.NewRPCFlightClient(conn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	hCustomer := customer_handler.NewCustomerHandler(customerClient)
	hBooking := booking_handler.NewBookingHandler(bookingClient)
	hFlight := flight_handler.NewFlightHandler(flightClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/customer/create", hCustomer.CreateCustomer)
	gr.POST("/booking/create", hBooking.CreateBooking)
	gr.POST("/flight/create", hFlight.CreateFlight)

	//Listen and serve
	http.ListenAndServe(":3333", g)
}

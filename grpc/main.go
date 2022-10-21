package main

import (
	"flag"
	"fmt"
	booking_handler "gin-tuns_go_flight/grpc/booking-grpc/handler"
	booking_repo "gin-tuns_go_flight/grpc/booking-grpc/repo"
	customer_handler "gin-tuns_go_flight/grpc/customer-grpc/handler"
	customer_repo "gin-tuns_go_flight/grpc/customer-grpc/repo"
	flight_handler "gin-tuns_go_flight/grpc/flight-grpc/handler"
	flight_repo "gin-tuns_go_flight/grpc/flight-grpc/repo"
	"gin-tuns_go_flight/helper"
	"gin-tuns_go_flight/intercepter"
	"gin-tuns_go_flight/pb"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configFile = flag.String("config-file", "../helper/config.yml", "Location of config file")
	port       = flag.Int("port", 2222, "Port of grpc")
)

func init() {
	flag.Parse()
}

func main() {
	err := helper.AutoBindConfig(*configFile)
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			intercepter.UnaryServerLoggingIntercepter(logger),
		)),
	)

	reflection.Register(s)

	// Initial customer repo START
	customerRepository, err := customer_repo.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := customer_handler.NewCustomerHandler(customerRepository)
	if err != nil {
		panic(err)
	}
	pb.RegisterRPCCustomerServer(s, h)
	// Initial customer repo END

	// Initial Flight repo START
	flightRepository, errFlight := flight_repo.NewDBManager()
	if errFlight != nil {
		panic(errFlight)
	}

	hFlight, errFlight := flight_handler.NewFlightHandler(flightRepository)
	if errFlight != nil {
		panic(errFlight)
	}
	pb.RegisterRPCFlightServer(s, hFlight)
	// Initial Flight repo END

	// Initial Booking repo START
	bookingRepository, errBooking := booking_repo.NewDBManager()
	if errBooking != nil {
		panic(errBooking)
	}

	hBooking, errBooking := booking_handler.NewBookingHandler(bookingRepository)
	if errBooking != nil {
		panic(errBooking)
	}
	pb.RegisterRPCBookingServer(s, hBooking)
	// Initial Booking repo END

	fmt.Printf("Listen at port: %v\n", *port)

	s.Serve(listen)
}

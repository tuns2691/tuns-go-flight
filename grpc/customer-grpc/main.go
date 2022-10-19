package main

import (
	"flag"
	"fmt"
	"gin-tuns_go_flight/grpc/customer-grpc/handler"
	"gin-tuns_go_flight/grpc/customer-grpc/repo"
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
	configFile = flag.String("config-file", "config.yml", "Location of config file")
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

	customerRepository, err := repo.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handler.NewCustomerHandler(customerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterRPCCustomerServer(s, h)

	fmt.Printf("Listen at port: %v\n", *port)

	s.Serve(listen)
}

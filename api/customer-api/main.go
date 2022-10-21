package main

import (
	customer_handler "gin-tuns_go_flight/api/customer-api/handler"
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
	customerConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	customerClient := pb.NewRPCCustomerClient(customerConn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	h := customer_handler.NewCustomerHandler(customerClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	//Create routes
	gr := g.Group("/v1/api")

	gr.POST("/create", h.CreateCustomer)

	//Listen and serve
	http.ListenAndServe(":3333", g)
}

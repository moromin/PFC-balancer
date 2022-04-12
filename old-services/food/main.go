package main

import (
	"fmt"
	"log"
	"net"

	"github.com/moromin/PFC-balancer/services/food/config"
	"github.com/moromin/PFC-balancer/services/food/db"
	"github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/food/server"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatal("Failed to listening:", err)
	}

	fmt.Println("Food Service on", c.Port)

	s := server.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	proto.RegisterFoodServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/moromin/PFC-balancer/services/recipe/config"
	"github.com/moromin/PFC-balancer/services/recipe/db"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
	"github.com/moromin/PFC-balancer/services/recipe/server"
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

	fmt.Println("Recipe Service on", c.Port)

	s := server.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	proto.RegisterRecipeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

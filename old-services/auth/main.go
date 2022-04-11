package main

import (
	"fmt"
	"log"
	"net"

	"github.com/moromin/PFC-balancer/services/auth/config"
	"github.com/moromin/PFC-balancer/services/auth/db"
	"github.com/moromin/PFC-balancer/services/auth/proto"
	"github.com/moromin/PFC-balancer/services/auth/server"
	"github.com/moromin/PFC-balancer/services/auth/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatal("failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-service",
		ExpirationHours: 24,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatal("failed to listening:", err)
	}

	fmt.Println("Auth service on", c.Port)

	s := server.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	proto.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}

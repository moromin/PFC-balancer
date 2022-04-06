package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/services/gateway/config"
	"github.com/moromin/PFC-balancer/services/gateway/pkg/auth"
	"github.com/moromin/PFC-balancer/services/gateway/pkg/food"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed at config")
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	food.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}

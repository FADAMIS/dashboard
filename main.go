package main

import (
	"github.com/FADAMIS/dashboard/api"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// participant register
	server.POST("/register", api.Register)

	// admin login
	server.POST("/login", api.Login)

	server.POST("/order/:name", api.OrderFood)

	server.POST("/food", api.AddFood)
	server.GET("/food", api.GetFoods)

	server.GET("/participants", api.GetParticipants)

	server.Run()
}

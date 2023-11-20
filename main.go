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
	server.POST("/admin/login", api.Login)

	server.POST("/order/:name", api.OrderFood)

	server.POST("/admin/food", api.AddFood)

	// returns all foods WITHOUT listed participants
	server.GET("/food", api.GetFoods)

	server.GET("/admin/participants", api.GetParticipants)

	// returns all foods WITH listed participants
	server.GET("/admin/food", api.GetFoodsAdmin)

	// uploaded images go here
	server.Static("/images", "./uploads")

	server.Run()
}

package main

import (
	"github.com/gin-contrib/gzip"

	"github.com/FADAMIS/dashboard/api"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(gzip.Gzip(gzip.DefaultCompression))

	// participant register
	server.POST("/api/register", api.Register)

	// admin login
	server.POST("/api/admin/login", api.Login)

	server.POST("/api/order/:name", api.OrderFood)

	server.POST("/api/admin/food", api.AddFood)

	// returns all foods WITHOUT listed participants
	server.GET("/api/food", api.GetFoods)

	server.GET("/api/admin/participants", api.GetParticipants)

	// returns all foods WITH listed participants
	server.GET("/api/admin/food", api.GetFoodsAdmin)

	// uploaded images go here
	server.Static("/images", "./uploads")

	server.Run()
}

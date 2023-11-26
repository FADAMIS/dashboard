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

	// returns all foods WITHOUT listed participants
	server.GET("/api/food", api.GetFoods)

	server.GET("/api/admin/participants", api.GetParticipants)

	// returns all foods WITH listed participants
	server.GET("/api/admin/food", api.GetFoodsAdmin)
	server.POST("/api/admin/food", api.AddFood)

	server.POST("/api/admin/camp", api.AddCamp)
	server.GET("/api/admin/camp", api.GetCampsAdmin)

	// disable registration and send participant list
	server.POST("/api/admin/process", api.ProcessCamp)

	server.GET("/api/camp", api.GetCamps)

	// uploaded images go here
	server.Static("/api/images", "./uploads")

	server.Run()
}

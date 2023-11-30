package main

import (
	"flag"
	"fmt"

	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"

	"github.com/FADAMIS/dashboard/api"
	"github.com/FADAMIS/dashboard/db"
	"github.com/gin-gonic/gin"
)

func main() {
	var isSetup bool
	flag.BoolVar(&isSetup, "config", false, "Initial setup")
	flag.Parse()

	if isSetup {
		Setup()

		return
	}

	db.InitDB()

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
	server.POST("/api/admin/close", api.CloseCamp)

	server.GET("/api/camp", api.GetCamps)

	// uploaded images go here
	server.Static("/api/images", "./uploads")

	server.Run()
}

func Setup() {
	fmt.Println("**Config flag is set. Begin initial setup**")

	postgresConfig, _ := godotenv.Read("../postgres/postgres.example.env")

	pgAdminConfig, _ := godotenv.Read("../postgres/pgadmin.example.env")
	finalPgAdminConfig := make(map[string]string)

	sampleBackendConfig, _ := godotenv.Read("./example.env")
	finalBackendConfig := make(map[string]string)
	for k := range sampleBackendConfig {
		var envValue string

		fmt.Print("Enter value for " + k + ": ")
		fmt.Scanln(&envValue)

		finalBackendConfig[k] = envValue

		if k == "POSTGRES_USER" || k == "POSTGRES_PASSWORD" {
			postgresConfig[k] = envValue
		}
	}

	for k := range pgAdminConfig {
		var envValue string

		fmt.Print("Enter value for " + k + ": ")
		fmt.Scanln(&envValue)

		finalPgAdminConfig[k] = envValue
	}

	godotenv.Write(finalBackendConfig, "./.env")
	godotenv.Write(postgresConfig, "../postgres/postgres.env")
	godotenv.Write(finalPgAdminConfig, "../postgres/pgadmin.env")
}

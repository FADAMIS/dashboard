package api

import (
	"net/http"

	"github.com/FADAMIS/dashboard/db"
	"github.com/FADAMIS/dashboard/entities"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var participant entities.Participant
	ctx.Bind(&participant)

	participant.FoodID = 1

	db.RegisterParticipant(participant)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "register successful",
	})
}

func GetParticipants(ctx *gin.Context) {
	participants := db.GetParticipants()

	ctx.JSON(http.StatusOK, gin.H{
		"participants": participants,
	})
}

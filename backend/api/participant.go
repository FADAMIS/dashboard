package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/FADAMIS/dashboard/db"
	"github.com/FADAMIS/dashboard/entities"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var participant entities.Participant
	ctx.Bind(&participant)

	participant.FoodID = 1

	participant.Name = strings.TrimSpace(participant.Name)
	participant.Surname = strings.TrimSpace(participant.Surname)

	check := false
	camps := db.GetCamps()
	for _, c := range camps {
		// if camp id exists and camp registration is not expired
		if c.ID == participant.CampID && !c.Closed {
			check = true
			db.RegisterParticipant(participant, c)
			SendRegisterConfirmation(participant.Email, participant.Name, participant.Surname, c.Name, c.Date)
			break
		}

		check = false
	}

	if !check {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Camp not available",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "register successful",
	})
}

func GetParticipants(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !IsSessionValid(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	participants := db.GetParticipants()

	ctx.JSON(http.StatusOK, gin.H{
		"participants": participants,
	})
}

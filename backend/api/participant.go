package api

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"slices"
	"strconv"
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
	participant.Email = strings.TrimSpace(participant.Email)
	participant.Phone = strings.TrimSpace(participant.Phone)

	_, err := mail.ParseAddress(participant.Email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "wrong email address",
		})

		return
	}

	participantFullInfo := participant.Name + participant.Surname + participant.Email + participant.Phone + strconv.Itoa(int(participant.CampID))
	isRegistered := slices.IndexFunc(db.GetParticipants(), func(p entities.Participant) bool {
		return p.Name+p.Surname+p.Email+p.Phone+strconv.Itoa(int(p.CampID)) == participantFullInfo
	})

	if isRegistered != -1 {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "participant already registered",
		})

		return
	}

	camps := db.GetCamps()

	// tells if camps exists
	campIndex := slices.IndexFunc(db.GetCamps(), func(c entities.Camp) bool {
		// if camp id exists and camp registration is not expired
		return c.ID == participant.CampID && !c.Closed
	})

	if campIndex == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "camp not available",
		})

		return
	}

	db.RegisterParticipant(participant, camps[campIndex])
	SendRegisterConfirmation(participant.Email, participant.Name, participant.Surname, camps[campIndex].Name, camps[campIndex].Date)

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

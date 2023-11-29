package api

import (
	"encoding/json"
	"net/http"

	"github.com/FADAMIS/dashboard/db"
	"github.com/FADAMIS/dashboard/entities"
	"github.com/gin-gonic/gin"
)

func AddCamp(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !IsSessionValid(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	var camp entities.Camp
	ctx.Bind(&camp)

	camp.Closed = false

	db.AddCamp(camp)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "camp added",
		"camp":    camp.Name,
	})
}

func RemoveCamp(camps []entities.Camp, index int) []entities.Camp {
	ret := make([]entities.Camp, 0)
	ret = append(ret, camps[:index]...)
	return append(ret, camps[index+1:]...)
}

func GetCamps(ctx *gin.Context) {
	camps := db.GetCamps()

	filteredCamps := camps

	// if registration is closed, remove camp from returned array
	for i, c := range camps {
		if c.Closed {
			filteredCamps = RemoveCamp(camps, i)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"camps": filteredCamps,
	})
}

func GetCampsAdmin(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !IsSessionValid(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	camps := db.GetCampsAdmin()

	ctx.JSON(http.StatusOK, gin.H{
		"camps": camps,
	})
}

// Disable registration and send participant list
func CloseCamp(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !IsSessionValid(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	var camp entities.Camp
	ctx.Bind(&camp)

	allCamps := db.GetCampsAdmin()

	contains := false
	for _, c := range allCamps {
		if camp.Name == c.Name && camp.ID == c.ID {
			camp = c
			contains = true
			break
		}

		contains = false
	}

	if !contains {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "camp not found",
		})

		return
	}

	if !camp.Closed {
		camp.Closed = true
		db.CloseCamp(camp)
		SendParticipantList(camp)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "camp closed",
			"camp":    camp.Name,
		})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "camp was already closed",
		})
	}
}

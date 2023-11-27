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

	camp.Processed = false

	db.AddCamp(camp)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "camp added",
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

	// if registration is closed, remove camp
	for i, c := range camps {
		if c.Processed {
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

	filteredCamps := camps

	// if registration is closed, remove camp
	for i, c := range camps {
		if c.Processed {
			filteredCamps = RemoveCamp(camps, i)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"camps": filteredCamps,
	})
}

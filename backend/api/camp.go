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

	db.AddCamp(camp)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "camp added",
	})
}

func GetCamps(ctx *gin.Context) {
	camps := db.GetCamps()

	ctx.JSON(http.StatusOK, gin.H{
		"camps": camps,
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

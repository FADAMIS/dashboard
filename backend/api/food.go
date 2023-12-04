package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/FADAMIS/dashboard/db"
	"github.com/FADAMIS/dashboard/entities"
	"github.com/gin-gonic/gin"
)

func OrderFood(ctx *gin.Context) {
	participantHash := ctx.Param("name")

	var food entities.Food
	ctx.Bind(&food)

	foods := db.GetFoods()
	foodIndex := slices.IndexFunc(foods, func(f entities.Food) bool { return f.Name == food.Name && f.ID == food.ID })
	if foodIndex == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "food not found",
		})

		return
	}

	food = foods[foodIndex]

	participants := db.GetParticipants()
	participantIndex := slices.IndexFunc(participants, func(p entities.Participant) bool {
		/*
			hash is combination of name, surname and camp ID
			camp ID because one participant can be registered to multiple camps
		*/
		preHash := p.Name + p.Surname + strconv.Itoa(int(p.CampID))
		sum := sha256.Sum256([]byte(preHash))
		hashed := hex.EncodeToString(sum[:])

		return participantHash == hashed
	})

	if participantIndex == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Participant not found",
		})

		return
	}

	db.OrderFood(participants[participantIndex], food)
	ctx.JSON(http.StatusOK, gin.H{
		"name": participants[participantIndex].Name + " " + participants[participantIndex].Surname,
		"food": food.Name,
	})
}

func AddFood(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !IsSessionValid(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	fileName := UploadImage(ctx)
	if fileName == "" {
		return
	}

	imagePath := "/api/images/" + fileName

	var food entities.Food
	ctx.Bind(&food)
	food.Imagepath = imagePath

	db.AddFood(food)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "food added",
		"food":    food.Name,
	})
}

func GetFoods(ctx *gin.Context) {
	foods := db.GetFoods()

	ctx.JSON(http.StatusOK, gin.H{
		"foods": foods,
	})
}

func GetFoodsAdmin(ctx *gin.Context) {
	var session entities.Session
	cookie, _ := ctx.Cookie("session")
	json.Unmarshal([]byte(cookie), &session)

	if !IsSessionValid(session) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})

		return
	}

	foods := db.GetFoodsAdmin()

	ctx.JSON(http.StatusOK, gin.H{
		"foods": foods,
	})
}

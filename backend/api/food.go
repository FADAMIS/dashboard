package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FADAMIS/dashboard/db"
	"github.com/FADAMIS/dashboard/entities"
	"github.com/gin-gonic/gin"
)

func OrderFood(ctx *gin.Context) {
	participantHash := ctx.Param("name")

	var food entities.Food
	ctx.Bind(&food)

	check := false
	foods := db.GetFoods()
	for _, f := range foods {

		if f.Name == food.Name && f.ID == food.ID {
			food = f
			check = true
			break
		}

		check = false
	}

	if !check {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Food not found",
		})

		return
	}

	participants := db.GetParticipants()
	for _, p := range participants {
		/*
			hash is combination of name, surname and camp ID
			because one participant can be registered to multiple camps
		*/
		preHash := p.Name + p.Surname + strconv.Itoa(int(p.CampID))
		sum := sha256.Sum256([]byte(preHash))
		hashed := hex.EncodeToString(sum[:])

		if participantHash == hashed {
			db.OrderFood(p, food)
			ctx.JSON(http.StatusOK, gin.H{
				"name": preHash,
				"food": food.Name,
			})

			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Participant not found",
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

	imagePath := "/images/" + fileName

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

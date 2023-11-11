package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/FADAMIS/dashboard/db"
	"github.com/FADAMIS/dashboard/entities"
	"github.com/gin-gonic/gin"
)

func OrderFood(ctx *gin.Context) {
	nameHash := ctx.Param("name")

	var food entities.Food
	ctx.Bind(&food)

	check := 0
	foods := db.GetFoods()
	for _, f := range foods {

		if f.Name == food.Name && f.ID == food.ID {
			break
		} else {
			check++
		}

		if check < len(foods) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Food not found",
			})

			return
		}
	}

	participants := db.GetParticipants()
	for _, p := range participants {
		fullName := p.Name + " " + p.Surname
		sum := sha256.Sum256([]byte(fullName))
		hashed := hex.EncodeToString(sum[:])
		if nameHash == hashed {
			db.OrderFood(p, food)
			ctx.JSON(http.StatusOK, gin.H{
				"name": fullName,
				"food": food.Name,
			})

			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

func AddFood(ctx *gin.Context) {
	var food entities.Food
	ctx.Bind(&food)

	db.AddFood(food)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "food added",
		"food":    food.Name,
	})
}

func GetFoods(ctx *gin.Context) {
	foods := db.GetFoods()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"foods":   foods,
	})
}

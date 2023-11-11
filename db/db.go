package db

import (
	"github.com/FADAMIS/dashboard/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dbInfo := "host=localhost user=fanda password=test123 dbname=dashboard port=5432 sslmode=disable TimeZone=Europe/Prague"

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Food{}, &entities.Participant{})

	return db, nil
}

var DB, _ = InitDB()

func GetDB() *gorm.DB {
	return DB
}

func AddFood(food entities.Food) {
	db := GetDB()

	db.Create(&food)
}

func GetFoods() []entities.Food {
	db := GetDB()

	var foods []entities.Food
	db.Model(&entities.Food{}).Preload("Participants").Find(&foods)

	return foods
}

func RegisterParticipant(participant entities.Participant) {
	db := GetDB()

	db.Create(&participant)
}

func GetParticipants() []entities.Participant {
	db := GetDB()

	var participants []entities.Participant
	db.Find(&participants)

	return participants
}

func OrderFood(participant entities.Participant, food entities.Food) {
	db := GetDB()
	food.Participants = append(food.Participants, participant)

	db.Save(&food)
}

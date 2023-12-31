package db

import (
	"os"

	"github.com/FADAMIS/dashboard/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	godotenv.Load()
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv(("POSTGRES_PASSWORD"))

	dbInfo := "host=database.docker user=" + dbUser + " password=" + dbPass + " dbname=dashboard port=5432 sslmode=disable TimeZone=Europe/Prague"

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Food{}, &entities.Camp{}, &entities.Participant{}, &entities.Session{}, &entities.Admin{})

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

// returns all foods WITH plisted articipants
func GetFoodsAdmin() []entities.Food {
	db := GetDB()

	var foods []entities.Food
	db.Model(&entities.Food{}).Preload("Participants").Find(&foods)

	return foods
}

// returns all foods WITHOUT listed participants
func GetFoods() []entities.Food {
	db := GetDB()

	var foods []entities.Food
	db.Find(&foods)

	return foods
}

func RegisterParticipant(participant entities.Participant, camp entities.Camp) {
	db := GetDB()
	camp.Participants = append(camp.Participants, participant)

	db.Save(&camp)
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

func AddSession(session entities.Session) {
	db := GetDB()

	db.Create(&session)
}

func GetSessions() []entities.Session {
	db := GetDB()

	var sessions []entities.Session
	db.Find(&sessions)

	return sessions
}

func DeleteSession(session entities.Session) {
	db := GetDB()

	db.Delete(&session)
}

func GetAdmins() []entities.Admin {
	db := GetDB()

	var admin []entities.Admin
	db.Find(&admin)

	return admin
}

func UpdateAdmin(admin entities.Admin) {
	db := GetDB()

	db.Save(&admin)

	// When admin updates his password, every session gets deleted
	db.Delete(&entities.Session{}, "username = ?", admin.Username)
}

func AddCamp(camp entities.Camp) {
	db := GetDB()

	db.Create(&camp)
}

func CloseCamp(camp entities.Camp) {
	db := GetDB()

	db.Model(&entities.Camp{}).Where("id = ?", camp.ID).Update("closed", true)
}

func GetCamps() []entities.Camp {
	db := GetDB()

	var camps []entities.Camp
	db.Find(&camps)

	return camps
}

func GetCampsAdmin() []entities.Camp {
	db := GetDB()

	var camps []entities.Camp
	db.Model(&entities.Camp{}).Preload("Participants").Find(&camps)

	return camps
}

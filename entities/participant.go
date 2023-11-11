package entities

type Participant struct {
	ID      uint
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	FoodID  uint   `json:"food_id"`
}

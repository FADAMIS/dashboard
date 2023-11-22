package entities

type Participant struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	FoodID  uint   `json:"food_id"`
	CampID  uint   `json:"camp_id"`
}

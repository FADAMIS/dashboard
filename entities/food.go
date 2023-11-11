package entities

type Food struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	Participants []Participant `json:"participants"`
}

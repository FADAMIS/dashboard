package entities

type Camp struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	Date         string        `json:"date"`
	Expires      int64         `json:"expires"`
	Participants []Participant `json:"participants"`
}

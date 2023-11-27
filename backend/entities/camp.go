package entities

type Camp struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	Date         int64         `json:"date"`
	Participants []Participant `json:"participants"`
	// tells if camp is closed - turns to "true" when email with listed participants is sent
	Closed bool `json:"processed"`
}

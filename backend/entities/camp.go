package entities

type Camp struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	Date         int64         `json:"date"`
	Participants []Participant `json:"participants"`
	// tells if camp is processed - turns to "true" when email with listed participants is sent
	Processed bool `json:"processed"`
}

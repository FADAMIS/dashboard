package entities

type Food struct {
	ID           uint          `json:"id" form:"id"`
	Name         string        `json:"name" form:"name"`
	Imagepath    string        `json:"image_path" form:"image_path"`
	Participants []Participant `json:"participants" form:"participants"`
}

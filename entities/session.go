package entities

type Session struct {
	UUID     string
	Username string
	Expires  int64
}

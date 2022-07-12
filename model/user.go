package model

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int64  `json:"points"`
}

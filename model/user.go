package model

// import "github.com/go-playground/validator/v10"

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int64  `json:"points" binding:"required,gte=18,lte=70"`
}

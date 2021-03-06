package model

import "gin_demo/utils"

type SignUpParam struct {
	Age        uint8        `json:"age" binding:"gte=1,lte=130"`
	Name       string       `json:"name" binding:"required"`
	Email      string       `json:"email" binding:"required,email"`
	Password   string       `json:"password" binding:"required"`
	RePassword string       `json:"re_password" binding:"required,eqfield=Password"`
	CreateTime utils.MyTime `json:"create_time" binding:"required,timing" time_format:"2006-01-02"`
}

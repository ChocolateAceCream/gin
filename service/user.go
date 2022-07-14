package service

import (
	"fmt"
	"gin_demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// return all users info.
// func (s *Service) GetAllUsers(c *gin.Context) (users []*model.User, err error) {
func (s *Service) GetAllUsers(c *gin.Context) (code int, obj interface{}) {
	fmt.Println("GetAllUsers svc", s)
	fmt.Println("GetAllUsers svc.Dao", s.Dao)

	users, err := s.Dao.FindAllUsers()

	if err != nil {
		code = http.StatusInternalServerError
		obj = gin.H{"status": "400", "errorMsg": err}
	} else {
		code = http.StatusOK
		obj = gin.H{"status": "200", "users": users}
	}
	return
}

func (s *Service) AddUser(c *gin.Context) (code int, obj interface{}) {
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("----------err: ", err)
		code = http.StatusInternalServerError
		obj = gin.H{"status": "400", "errorMsg": err}
		return
	}
	fmt.Println("----------user: ", *user)

	//TODO: verification and validation
	id, err := s.Dao.AddUser(user)
	if err != nil {
		code = http.StatusInternalServerError
		obj = gin.H{"status": "400", "errorMsg": err}
	} else {
		code = http.StatusOK
		obj = gin.H{"status": "200", "id": id}
	}
	return
}

func (s *Service) EditUser(c *gin.Context) (code int, obj interface{}) {
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("----------err: ", err)
		code = http.StatusInternalServerError
		obj = gin.H{"status": "400", "errorMsg": err}
		return
	}
	fmt.Println("----------user: ", *user)

	//TODO: verification and validation
	count, err := s.Dao.EditUser(user)
	if err != nil {
		code = http.StatusInternalServerError
		obj = gin.H{"status": "400", "errorMsg": err}
	} else {
		code = http.StatusOK
		obj = gin.H{"status": "200", "update succ:": count}
	}
	return
}

func (s *Service) DeleteUser(c *gin.Context) (code int, obj interface{}) {
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("----------err: ", err)
		code = http.StatusInternalServerError
		obj = gin.H{"status": "400", "errorMsg": err}
		return
	}
	fmt.Println("----------user: ", *user)

	//TODO: verification and validation
	count, err := s.Dao.DeleteUser(user)
	if err != nil {
		code = http.StatusInternalServerError
		obj = gin.H{"status": "400", "errorMsg": err}
	} else {
		code = http.StatusOK
		obj = gin.H{"status": "200", "update succ:": count}
	}
	return
}

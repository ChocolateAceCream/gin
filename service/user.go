package service

import (
	"fmt"
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

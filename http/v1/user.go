package v1

import (
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	// var parsedQuery LoginQuery
	// id := c.Param("id")
	// fmt.Println("----- id----", id)
	// if err := c.BindQuery(&parsedQuery); err != nil {
	// 	fmt.Println("----- parsedQuery----", parsedQuery)
	// 	c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
	// }
	// if parsedQuery.User != "admin" || parsedQuery.Password != "123qwe" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errorMsg": "not authorized"})
	// 	return
	// }
	c.JSON(svc.GetAllUsers(c))

	// c.JSON(http.StatusOK, gin.H{"status": "200", "from": "getinfo"})
}

func AddUser(c *gin.Context) {
	c.JSON(svc.AddUser(c))
}

func EditUser(c *gin.Context) {
	c.JSON(svc.EditUser(c))
}

func DeleteUser(c *gin.Context) {
	c.JSON(svc.DeleteUser(c))
}

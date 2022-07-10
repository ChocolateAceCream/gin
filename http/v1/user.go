package v1

import (
	"net/http"

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
	c.JSON(http.StatusOK, gin.H{"status": "200", "from": "getinfo"})
}

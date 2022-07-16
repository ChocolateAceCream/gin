package v1

import (
	"github.com/gin-gonic/gin"
)

// @Tags Base/v1
// @Summary Get all user
// @Produce  application/json
// @Success 200 {string} string "{"status":"200","users":[{"id":0,"firstName":"Ines","lastName":"Brushfield","points":947},{"id":0,"firstName":"Freddi","lastName":"Boagey","points":2967},{"id":0,"firstName":"Ambur","lastName":"Roseburgh","points":457},{"id":0,"firstName":"Clemmie","lastName":"Betchley","points":3675},{"id":0,"firstName":"Elka","lastName":"Twiddell","points":3073},{"id":0,"firstName":"Ilene","lastName":"Dowson","points":1672},{"id":0,"firstName":"Thacher","lastName":"Naseby","points":205},{"id":0,"firstName":"Romola","lastName":"Rumgay","points":1486},{"id":0,"firstName":"Levy","lastName":"Mynett","points":796},{"id":0,"firstName":"admin","lastName":"123qwe","points":54},{"id":0,"firstName":"admin","lastName":"123qwe","points":54},{"id":0,"firstName":"admin","lastName":"123qwe","points":54},{"id":0,"firstName":"Timothy","lastName":"White","points":96}]}"
// @Router /v1/customer/query [get]
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

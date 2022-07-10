package v1

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UrlHandler(c *gin.Context) {
	// name := c.Query("name") //if no query passed, then return empty string
	name := c.Query("name")
	// name := c.DefaultQuery("name", "default_name")
	msg, flag := c.Get("timestamp")
	if !flag {
		fmt.Println("error: ", flag)
	}
	fmt.Println("msg from middleWare: ", msg)
	var resp struct {
		Name string
		Time string
	}
	resp.Name = name
	resp.Time = time.Now().String()
	c.JSON(http.StatusOK, resp)
}

func ApiHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	//截取/
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
}

func FormHandler(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password ")
	c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, types: %s", username, password, types))
}

package v1

import (
	"fmt"
	"gin_demo/library/logger"
	"gin_demo/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func ValidationDemo(c *gin.Context) {
	var u *model.SignUpParam
	if err := c.ShouldBind(&u); err != nil {
		logger.ZapLog_V1.Error("ValidationDemo error", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// other logics

	c.JSON(http.StatusOK, "success")
}

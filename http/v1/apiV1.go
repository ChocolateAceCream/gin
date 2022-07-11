package v1

import (
	"gin_demo/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// load api v1 endpoints
func RouteLoader(r *gin.Engine) {
	v1 := r.Group("/v1")
	// {} is standard format for route group
	{
		v1.GET("/:name/*action", ApiHandler)
		v1.GET("/user", UrlHandler)
		v1.POST("/login/", Login)
		v1.GET("/info/:id", GetInfo)
		v1.GET("/session", middleware.SessionMiddleware(), SessionDemo)
		v1.GET("/lockDemo", LockDemo)

		// redirect request
		v1.GET("/baidu", func(c *gin.Context) {
			// http.StatusMovedPermanently: 301 Moved Permanently
			c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		})
		user := v1.Group("/customer")
		{
			user.GET("/query", GetAllUsers)
		}
	}
}

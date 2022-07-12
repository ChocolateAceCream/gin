package http

import (
	"gin_demo/http/middleware"
	v1 "gin_demo/http/v1"
	v2 "gin_demo/http/v2"

	"gin_demo/service"

	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service
)

func Init(s *service.Service) {
	svc = s
}

// *gin.Engine is the thing that can be passed around
func SetupRouter(r *gin.Engine) {
	// 20 times (8 times 2 ), 2^10 is 1024, so this means 8MB
	maxSize := 8 << 20

	// when uploading files, limit the max memory usage to 8MB (not the max file size)
	r.MaxMultipartMemory = int64(maxSize)

	// apply global middle ware
	r.Use(middleware.GetMiddleware()...)

	// optional you can apply local middle ware for certain endpoint
	// v2.POST("/form", MiddleWare(), formHandler)

	//load router endpoints, basically wrapper over the code block below
	v1.RouteLoader(r, svc)
	v2.RouteLoader(r, svc)
}

package v2

import (
	"gin_demo/service"

	"github.com/gin-gonic/gin"
)

var (
	svc *service.Service
)

func RouteLoader(r *gin.Engine, s *service.Service) {
	svc = s
	v2 := r.Group("/v2")
	{
		v2.POST("/form", FormHandler)
		v2.POST("/uploader", FileUploader)
		v2.POST("/uploaderV2", FileUploaderV2)
		v2.POST("/multiFilesUploader", MultiFilesUploader)

	}
}

package v2

import (
	"github.com/gin-gonic/gin"
)

func RouteLoader(r *gin.Engine) {
	v2 := r.Group("/v2")
	{
		v2.POST("/form", FormHandler)
		v2.POST("/uploader", FileUploader)
		v2.POST("/uploaderV2", FileUploaderV2)
		v2.POST("/multiFilesUploader", MultiFilesUploader)

	}
}

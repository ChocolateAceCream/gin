package v2

import (
	"fmt"
	"gin_demo/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FormHandler(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password ")
	c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, types: %s", username, password, types))
}

func FileUploader(c *gin.Context) {
	//fileKey is the key that passed in form-data
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, fmt.Sprint("upload failed", err))
	}
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
}

func FileUploaderV2(c *gin.Context) {
	_, headers, err := c.Request.FormFile("file")
	if err != nil {
		c.String(500, fmt.Sprint("upload failed", err))
	}

	contentType := headers.Header.Get("Content-Type")
	fmt.Println(contentType)
	// if contentType != "text/plain" {
	allowedTypes := []string{"application/gzip", "image/jpeg", "text/plain"}

	// check uploaded file type is allowed or not
	if !utils.Contains(allowedTypes, contentType) {
		fmt.Printf("content type %v is not allowed", contentType)
		return
	}

	// check uploaded file size
	fmt.Println("----file size---", headers.Size)
	if headers.Size > 1024*1024*2 {
		fmt.Println("file size exceed limit, should below 2MB")
		return
	}

	timer := time.Now().Format("15:04:05")
	c.SaveUploadedFile(headers, timer+headers.Filename)
	c.String(http.StatusOK, headers.Filename)
}

func MultiFilesUploader(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}
	files := form.File["file"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}
	}
	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
}

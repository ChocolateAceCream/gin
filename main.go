package main

import (
	"fmt"
	"gin_demo/http"
	"io"
	"log"
	"os"

	"gin_demo/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() *gin.Engine {
	//load env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//setup log
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// make sure gin.Default() is executed after init logger
	r := gin.Default()

	//init service
	s := service.New()
	http.Init(s)

	//init router
	http.SetupRouter(r)
	// init redis session store
	// services.SetupSession(r)
	return r
}

func main() {

	r := Init()

	if err := r.Run("localhost:3000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}

	// TODO: add graceful shutdown, and close db using s.Close()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	// for {
	// 	si := <-c
	// 	log.Printf("get a signal %s", si.String())
	// 	switch si {
	// 	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	// 		log.Println(" exit")
	// 		// s.Close()
	// 		time.Sleep(10 * time.Second)
	// 		return
	// 	case syscall.SIGHUP:
	// 	default:
	// 		return
	// 	}
	// }
}

package main

import (
	"fmt"
	"gin_demo/global"
	"gin_demo/http"

	"gin_demo/library/logger"
	"gin_demo/service"
	"gin_demo/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() *gin.Engine {

	// setup zap logger
	logger.Init()

	//load env variable
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//load config.yaml
	err = godotenv.Load("config.yaml")
	if err != nil {
		panic(err)
	}

	// init viper
	global.Config = utils.Viper("config.yaml")

	// e.g usage:
	// jwt := global.Config.GetStringMap("jwt")
	// fmt.Println("----jwt----", jwt)

	// make sure gin.Default() is executed after init logger
	r := gin.Default()

	//register validator
	utils.InitValidator()

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

	if err := r.Run(); err != nil {
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

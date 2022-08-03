package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// TODO:
// 优先级: 命令行 > 环境变量 > 默认值

func Viper(path ...string) *viper.Viper {
	var config string

	config = path[0]
	fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		// 	fmt.Println(err)
		// }
	})
	// if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
	// 	fmt.Println(err)
	// }

	// // root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	// global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	// global.BlackCache = local_cache.NewCache(
	// 	local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
	// )
	return v
}

package main

import (
	"log"
	"www.github.com/ygxiaobai111/qiniu/config"
	"www.github.com/ygxiaobai111/qiniu/routes"
)

func main() {
	//配置信息初始化
	err := config.Init()
	if err != nil {
		log.Println(err)

	}
	r := routes.NewRouter()
	err = r.Run(":8811")
	if err != nil {
		log.Println(err)

	}
}

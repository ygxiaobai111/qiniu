package main

import (
	"log"
	"www.github.com/ygxiaobai111/qiniu/server/config"
	"www.github.com/ygxiaobai111/qiniu/server/routes"
)

//	@title			视频 API
//	@version		1.0
//	@description	This is cxy api docs.
//	@license.name	Apache 2.0
//	@contact.name
//	@contact.url	https://github.com/swaggo/swag/blob/master/README_zh-CN.md
//	@host			localhost:8811
//	@BasePath		/
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

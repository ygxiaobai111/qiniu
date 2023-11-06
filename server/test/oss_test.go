package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
	"www.github.com/ygxiaobai111/qiniu/server/repository/oss"
)

func oss_test(m *testing.M) {
	oss.Init()
	r := gin.Default()
	r.POST("/p", func(context *gin.Context) {
		file, err := context.FormFile("data")
		log.Println("err:", err)
		fileF, err := file.Open()
		log.Println("err:", err)
		f := make([]byte, file.Size)

		fileF.Read(f)
		log.Println("f:", len(f))
		vname, _ := oss.AddVideo(100001, "testTitle", f)
		fmt.Println("name:", vname)
	})
	r.Run(":8080")
	m.Run()
}

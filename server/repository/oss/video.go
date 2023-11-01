package oss

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"time"
)

func md5digest(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x_%d", hash, time.Now().Unix())
	return md5str
}

func Add(authorId int, title string, data []byte) (string, error) {

	digest := md5digest(title)
	log.Println("1", len(data))
	// 生成视频和图片的Key
	videoKey := fmt.Sprintf("public/%d/%s_%s.mp4", authorId, title, digest)
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:wmText":      "你牛",
			"x:wmGravity":   "NorthWest",
			"x:wmFontColor": "FFFFFF",
			"x:wmDissolve":  "100",
			"x:wmFontSize":  "200",
		},
	}
	// 调用PutFile方法上传文件
	err := FormUploader.Put(context.Background(), &Ret, *UpToken, videoKey, bytes.NewReader(data), int64(len(data)), &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	videoUrl := MYURL + Ret.Key
	//coverUrl := videoUrl + "?vframe/jpg/offset/1"
	return videoUrl, nil
}

package oss

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"log"
)

func md5digest(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)
	return md5str
}

func Add(authorId int, title string, data []byte) (string, error) {

	digest := md5digest(title)
	log.Println("1", len(data))
	// 生成视频和图片的Key
	videoKey := fmt.Sprintf("public/%d/%s_%s.mp4", authorId, title, digest)

	// 调用PutFile方法上传文件
	err := FormUploader.Put(context.Background(), &Ret, *UpToken, videoKey, bytes.NewReader(data), int64(len(data)), nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	videoUrl := "http://http://www.xzkckj.cn/" + Ret.Key
	//coverUrl := videoUrl + "?vframe/jpg/offset/1"
	return videoUrl, nil
}

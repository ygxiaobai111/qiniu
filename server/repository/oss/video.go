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

<<<<<<< HEAD
func AddWatermark(videoBytes []byte, watermarkText string) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(videoBytes))
	if err != nil {
		return nil, err
	}

	watermarkedImg := imaging.Clone(img)

	drawer := &font.Drawer{
		Dst:  watermarkedImg,
		Src:  image.White,
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{},
	}
	drawer.DrawString(watermarkText)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, watermarkedImg)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
func AddVideo(authorId uint, title string, data []byte) (string, error) {
=======
func AddVideo(authorId int, title string, data []byte) (string, error) {
>>>>>>> c29129335107c819911416d59d6b4f0f9a9fcff0

	digest := md5digest(title)
	log.Println("1", len(data))
	// 生成视频和图片的Key
	videoKey := fmt.Sprintf("public/%d/%s_%s.mp4", authorId, title, digest)
	PubKey := fmt.Sprintf("%s_%s.mp4", title, digest)
	ret := storage.PutRet{}
	// 调用PutFile方法上传文件
	err := FormUploader.Put(context.Background(), &ret, *UpToken, videoKey, bytes.NewReader(data), int64(len(data)), nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	videoUrl := MYURL + "shuiyin/" + PubKey

	//coverUrl := videoUrl + "?vframe/jpg/offset/1"
	return videoUrl, nil
}
func AddImage(authorId uint, title string, data []byte) (string, error) {

	digest := md5digest(title)
	log.Println("1", len(data))
	// 生成视频和图片的Key
	imageKey := fmt.Sprintf("image/%d/%s_%s.jpg", authorId, title, digest)

	ret := storage.PutRet{}
	// 调用PutFile方法上传文件
	err := FormUploader.Put(context.Background(), &ret, *UpToken, imageKey, bytes.NewReader(data), int64(len(data)), nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	imageUrl := MYURL + imageKey

	//coverUrl := videoUrl + "?vframe/jpg/offset/1"
	return imageUrl, nil
}

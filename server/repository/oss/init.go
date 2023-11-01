package oss

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"www.github.com/ygxiaobai111/qiniu/server/config"
)

var Viper *viper.Viper
var UpToken *string
var FormUploader *storage.FormUploader
var Ret storage.PutRet

// 我的资源域名
var MYURL string

// Init 传入包含AK和SK的文件路径,必须为yml文件 格式为qiniu.accessKey和qiniu.secretKey
func Init() (err error) {

	accessKey := config.AccessKey
	secretKey := config.SecretKey
	MYURL = config.MYURL
	// 设置要上传的空间
	bucket := config.Bucket

	// 生成上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	token := putPolicy.UploadToken(mac)
	UpToken = &token

	// 配置上传参数
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan // 设置存储区域
	cfg.UseHTTPS = false           // 是否使用https域名
	cfg.UseCdnDomains = false      // 是否使用cdn加速域名

	// 构建表单上传的对象
	FormUploader = storage.NewFormUploader(&cfg)
	Ret = storage.PutRet{}
	return
}

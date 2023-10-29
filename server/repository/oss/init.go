package oss

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

var Viper *viper.Viper
var UpToken *string
var FormUploader *storage.FormUploader
var Ret storage.PutRet

// Init 传入包含AK和SK的文件路径,必须为yml文件 格式为qiniu.accessKey和qiniu.secretKey
func Init() {
	// 设置七牛云账号的AK和SK
	accessKey := "wccY3Xc1dpW5hGTYvvyYF4LnJZ1Rsk6-6mbbcxw_"
	secretKey := "P11e3W-EyhiFESh0XPRhQT7Oihvj6OO48mcVig1E"

	// 设置要上传的空间
	bucket := "dydemo-01"
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
}

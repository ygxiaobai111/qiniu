package config

import (
	"github.com/spf13/viper"
)

var (
	ServiceName string
	ServerIp    string
	ServerPort  uint64
	NacosIp     string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName int64

	AccessKey string
	SecretKey string
	MYURL     string
	Bucket    string
)

// nacos配置项
var (
	NacosAddress string
	NacosPort    uint64
)

// 微服务的服务名
var (
	VideoCenterServiceName string
)

// Init 初始化配置文件与引擎
func Init() error {

	// 设置配置文件的名称和路径
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// 获取配置项的值并赋值给变量
	DbHost = viper.GetString("mysql.DbHost")
	DbPort = viper.GetString("mysql.DbPort")
	DbUser = viper.GetString("mysql.DbUser")
	DbPassword = viper.GetString("mysql.DbPassword")
	DbName = viper.GetString("mysql.DbName")
	RedisDb = viper.GetString("redis.RedisDb")
	RedisAddr = viper.GetString("redis.RedisAddr")
	RedisPw = viper.GetString("redis.RedisPw")
	RedisDbName = viper.GetInt64("redis.RedisDbName")
	// 设置七牛云账号的AK和SK
	AccessKey = viper.GetString("qiniu.accessKey")
	SecretKey = viper.GetString("qiniu.secretKey")
	MYURL = viper.GetString("qiniu.url")
	// 设置要上传的空间
	Bucket = viper.GetString("qiniu.bucket")
	// mysql 连接地址

	return nil
}

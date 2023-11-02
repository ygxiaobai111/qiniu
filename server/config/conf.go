package config

import (
	"github.com/spf13/viper"
)

var (
	ServiceName string
	ServerIp    string
	ServerPort  string
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
	ServerPort = viper.GetString("application.Port")
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
	AccessKey = viper.GetString("qiniu.AccessKey")
	SecretKey = viper.GetString("qiniu.SecretKey")
	MYURL = viper.GetString("qiniu.Url")
	// 设置要上传的空间
	Bucket = viper.GetString("qiniu.Bucket")
	// mysql 连接地址

	return nil
}

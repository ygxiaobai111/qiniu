package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"math/rand"
	"time"
	"www.github.com/ygxiaobai111/qiniu/server/config"
)

var (
	EsClient   *elastic.Client
	globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func Init() (err error) {
	// 创建客户端
	EsClient, err = elastic.NewClient(elastic.SetURL(config.EsUrl), elastic.SetBasicAuth(config.EsUsername, config.EsPassword), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		return err
	}

	// 使用Ping方法检查ES是否可用
	info, code, err := EsClient.Ping(config.EsUrl).Do(context.Background())
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// 获取ES版本号
	esversion, err := EsClient.ElasticsearchVersion(config.EsUrl)
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	return nil
}

package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"math/rand"
	"time"
)

var (
	EsClient   *elastic.Client
	globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func Init() (err error) {
	// 创建客户端
	EsClient, err = elastic.NewClient(elastic.SetURL("http://8.130.100.107:9200"), elastic.SetBasicAuth("elastic", "123456"), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		return err
	}

	// 使用Ping方法检查ES是否可用
	info, code, err := EsClient.Ping("http://8.130.100.107:9200").Do(context.Background())
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// 获取ES版本号
	esversion, err := EsClient.ElasticsearchVersion("http://8.130.100.107:9200")
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	return nil
}

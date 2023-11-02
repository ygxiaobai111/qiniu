package indexs

import (
	"context"
	"www.github.com/ygxiaobai111/qiniu/server/repository/es"
)

// ExistsIndex (索引名)判断索引是否存在
func ExistsIndex(index string) bool {
	exists, _ := es.EsClient.IndexExists(index).Do(context.Background())
	return exists
}

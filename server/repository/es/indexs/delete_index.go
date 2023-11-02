package indexs

import (
	"context"
	"www.github.com/ygxiaobai111/qiniu/server/repository/es"

	"fmt"
)

// DeleteIndex (索引名)删除索引
func DeleteIndex(index string) {
	_, err := es.EsClient.
		DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(index, "索引删除成功")
}

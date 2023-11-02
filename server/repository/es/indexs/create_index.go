package indexs

import (
	"context"
	"es_study/global"
	"es_study/models"
	"fmt"
)

// CreateIndex 创建索引
// ExistsIndex判断索引是否存在
// DeleteIndex先删除再添加

func CreateIndex() {
	index := "video_index"
	if ExistsIndex(index) {
		// 索引存在，先删除，在创建
		DeleteIndex(index)
	}

	createIndex, err := global.ESClient.
		CreateIndex(index).
		BodyString(models.VideoModel{}.VideoMapping()).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(createIndex.Index, "索引创建成功")
}

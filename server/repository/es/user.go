package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/repository/es/models"
)

// UserCreate 向索引表传入用户id和用户名字
func UserCreate(id uint, name string) (err error) {
	user := models.UserModel{
		UserId:   id,
		NickName: name,
	}

	indexResponse, err := EsClient.Index().Index(user.UserIndex()).BodyJson(user).Do(context.Background())
	if err != nil {
		util.LogrusObj.Error("es_create_error:", err)
		return
	}
	util.LogrusObj.Printf("es_create_req:[id: %#v,name:%#v ],resp:[%#v]", id, name, indexResponse)
	return nil
}

// UserRetrieve 传入页码，一页显示条数，用户名，返回用户id数组
func UserRetrieve(page int, limit int, name string) ([]uint, error) {
	if limit == 0 {
		limit = 15
	}
	if page <= 0 {
		page = 1
	}
	from := (page - 1) * limit
	query := elastic.NewMatchQuery("nick_name", name)
	util.LogrusObj.Printf("Retrieving users: page=%d, limit=%d, name=%s", page, limit, name)
	res, err := EsClient.Search(models.UserModel{}.UserIndex()).Query(query).From(from).Size(limit).Do(context.Background())
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	var ids []uint
	for _, hit := range res.Hits.Hits {
		var source map[string]interface{}
		err = json.Unmarshal(hit.Source, &source)
		if err != nil {
			return nil, err
		}
		if id, ok := source["user_id"].(float64); ok {
			ids = append(ids, uint(id))
		} else {
			util.LogrusObj.Println("Warning: could not convert user_id to float64")
		}
	}
	return ids, nil
}

// UserUpdate 输入用户id和用户修改的姓名
func UserUpdate(id uint, nickName string) {
	query := elastic.NewTermQuery("user_id", id)
	script := elastic.NewScriptInline("ctx._source.nick_name = params.nick_name").Param("nick_name", nickName)
	updateByQuery := EsClient.UpdateByQuery().Index(models.UserModel{}.UserIndex()).Query(query).Script(script)
	res, err := updateByQuery.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	util.LogrusObj.Printf("%#v\n", res)
}

// UserDelete 输入用户id删除用户索引行记录
func UserDelete(id uint) {
	query := elastic.NewTermQuery("user_id", id)
	deleteByQuery := EsClient.DeleteByQuery().Index(models.UserModel{}.UserIndex()).Query(query)
	deleteResponse, err := deleteByQuery.Do(context.Background())
	if err != nil {
		util.LogrusObj.Println(err)
		return
	}
	util.LogrusObj.Println(deleteResponse)
}

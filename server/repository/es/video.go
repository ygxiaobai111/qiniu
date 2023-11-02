package es

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/repository/es/models"
)

// VideoCreate 向视频索引表传入用户id，视频id，标签id，视频标题
func VideoCreate(uid int64, vid int64, tid int64, name string) {
	user := models.VideoModel{
		UserId:         uid,
		VideoId:        vid,
		TagId:          tid,
		NickVideoTitle: name,
	}

	indexResponse, err := EsClient.Index().Index(user.VideoIndex()).BodyJson(user).Do(context.Background())
	if err != nil {
		util.LogrusObj.Println(err)
		return
	}
	util.LogrusObj.Printf("%#v\n", indexResponse.Id)
}

// VideoTitleRetrieve 页码，页显示条数，视频标题
func VideoTitleRetrieve(page int, limit int, name string) ([]int64, error) {
	from := limit * (page - 1)
	query := elastic.NewMatchQuery(`nick_video_title`, name)
	util.LogrusObj.Printf("Retrieving video titles: page=%d, limit=%d, name=%s\n", page, limit, name)
	var res, err = EsClient.Search(models.VideoModel{}.VideoIndex()).Query(query).From(from).Size(limit).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, hit := range res.Hits.Hits {
		var source map[string]interface{}
		err := json.Unmarshal(hit.Source, &source)
		if err != nil {
			return nil, err
		}
		if id, ok := source["video_id"].(float64); ok {
			ids = append(ids, int64(id))
		} else {
			util.LogrusObj.Println("Warning: could not convert video_id to float64")
		}
	}
	return ids, nil
}

// VideoTagRetrieve 页码，页显示条数，视频标签，是否随机
func VideoTagRetrieve(page int, limit int, Tag int64, shuffle bool) ([]int64, error) {
	from := limit * (page - 1)
	query := elastic.NewMatchQuery(`tag_id`, Tag)
	util.LogrusObj.Printf("Retrieving video tags: page=%d, limit=%d, tag=%d, shuffle=%v\n", page, limit, Tag, shuffle)
	var res, err = EsClient.Search(models.VideoModel{}.VideoIndex()).Query(query).From(from).Size(limit).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, hit := range res.Hits.Hits {
		var source map[string]interface{}
		err := json.Unmarshal(hit.Source, &source)
		if err != nil {
			return nil, err
		}
		if id, ok := source["video_id"].(float64); ok {
			ids = append(ids, int64(id))
		}
	}

	if shuffle {
		globalRand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
		util.LogrusObj.Println("Shuffled video IDs")
	}

	return ids, nil
}

// VideoUpdate 输入视频id，修改的标题
func VideoUpdate(id int64, nickVideoTitle string) {
	query := elastic.NewTermQuery("video_id", id)
	script := elastic.NewScriptInline("ctx._source.nick_video_title = params.nick_video_title").Param("nick_video_title", nickVideoTitle)
	updateByQuery := EsClient.UpdateByQuery().Index(models.VideoModel{}.VideoIndex()).Query(query).Script(script)
	res, err := updateByQuery.Do(context.Background())
	if err != nil {
		util.LogrusObj.Println(err)
		return
	}
	util.LogrusObj.Printf("%#v\n", res)
}

// VideoDelete 删除视频记录
func VideoDelete(id int64) {
	query := elastic.NewTermQuery("video_id", id)
	deleteByQuery := EsClient.DeleteByQuery().Index(models.VideoModel{}.VideoIndex()).Query(query)
	deleteResponse, err := deleteByQuery.Do(context.Background())
	if err != nil {
		util.LogrusObj.Println(err)
		return
	}
	util.LogrusObj.Println(deleteResponse)
}

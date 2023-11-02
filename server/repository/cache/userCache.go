package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// 维护用户画像的函数
func UpdateUserProfile(ctx context.Context, userKey string, tagID int64, score float64) error {
	// 设置用户画像中的标签分数
	z := redis.Z{
		Score:  score,
		Member: tagID,
	}
	_, err := RedisClient.ZAdd(ctx, "user:"+userKey, z).Result()

	if err != nil {

		return err
	}

	return nil
}

func GetTopTags(ctx context.Context, userKey string) []int64 {
	// 获取用户画像中的所有标签分数，并按分数排序
	scores, err := RedisClient.ZRangeWithScores(ctx, "user:"+userKey, 0, -1).Result()
	if err != nil {
		fmt.Println("Error getting user profile:", err)
		return nil
	}

	// 提取排名前三的标签ID和分数
	topTags := make([]int64, 0)
	for i := 0; i < 3 && i < len(scores); i++ {

		tagID := scores[i].Member.(int64)
		// 将标签ID转换为int64类型并添加到结果列表中
		topTags = append(topTags, tagID)
	}
	return topTags
}

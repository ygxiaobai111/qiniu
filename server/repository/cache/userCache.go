package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
)

// 维护用户画像的函数
func UpdateUserProfile(ctx context.Context, userKey string, tagID uint, score float64) error {
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

// IsFollow  是否关注
func IsFollow(ctx context.Context, uId, followId uint) bool {
	return RedisClient.
		Exists(ctx, GenFollowUserCacheKey(uId, followId)).
		Val() == 1
}

// AddFollow 关注关系缓存
func AddFollow(ctx context.Context, uId, followId uint) error {
	return RedisClient.
		Set(ctx, GenFollowUserCacheKey(uId, followId), 1, 7*24*time.Hour).
		Err()
}

// DeleteFollow 取关关系缓存
func DeleteFollow(ctx context.Context, uId, followId uint) error {
	return RedisClient.
		Del(ctx, GenFollowUserCacheKey(uId, followId)).
		Err()
}

// AddUser 用户信息缓存
func AddUser(ctx context.Context, uId uint, m map[string]interface{}) error {
	err := RedisClient.HSet(ctx, GenUserInfoCacheKey(uId), m).Err()
	if err != nil {
		util.LogrusObj.Error("redis_AddUser_error:", err)
		return err
	}
	// 设置键的过期时间为10分钟
	return RedisClient.Expire(ctx, GenUserInfoCacheKey(uId), 600*time.Second).Err()
}

// HasUser 判断Redis中是否存在某个用户信息缓存
func HasUser(ctx context.Context, uId uint) (cacheData map[string]string, err error) {
	// 获取一个哈希表中的所有字段和值
	cacheData, err = RedisClient.HGetAll(ctx, GenUserInfoCacheKey(uId)).Result()
	if err != nil {
		util.LogrusObj.Error("redis_HasUser_error:", err)
		return nil, err
	}
	// 判断用户是否存在，如果哈希表为空，表示不存在，如果不为空，表示存在
	return cacheData, nil
}

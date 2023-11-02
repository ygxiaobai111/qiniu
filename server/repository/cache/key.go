package cache

import (
	"strconv"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
)

const (
	// RankKey 每日排名
	RankKey             = "rank"
	SkillProductKey     = "skill:product:%d"
	SkillProductListKey = "skill:product_list"
	SkillProductUserKey = "skill:user:%s"
)

func PersonasKey(id uint) string {
	struId := strconv.Itoa(int(id))

	return "personas" + struId
}
func GenFollowUserCacheKey(userId, followUserId uint) string {
	return "follow:" + util.UintToStr(userId) + ":" + util.UintToStr(followUserId)
}

func GenUserInfoCacheKey(userId uint) string {
	return "user:info:" + util.UintToStr(userId)
}

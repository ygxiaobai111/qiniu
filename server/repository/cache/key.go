package cache

import (
	"strconv"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
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

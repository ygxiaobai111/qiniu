package cache

import (
	"strconv"
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

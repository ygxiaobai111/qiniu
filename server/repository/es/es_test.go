package es

import (
	"testing"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
)

func TestMain(m *testing.M) {
	Init()

	users, _ := UserRetrieve(0, 0, "天才")
	util.LogrusObj.Info(users)
}

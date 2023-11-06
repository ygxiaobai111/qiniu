package test

import (
	"testing"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/repository/es"
)

func TestMain(m *testing.M) {
	es.Init()

	users, _ := es.UserRetrieve(0, 0, "天才")
	util.LogrusObj.Info(users)
}

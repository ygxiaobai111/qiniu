package service

import (
	"context"
	"fmt"
	"time"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
)

func init() {
	//每过1天更新一次热门视频榜
	ticker := time.NewTicker(24 * time.Hour)

	go func() {
		for {
			select {
			case <-ticker.C:
				err := HotVideo(context.Background())
				util.LogrusObj.Error("HotVideoInitErr:", err)
			}
		}
	}()

	time.Sleep(100 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

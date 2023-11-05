package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

// AddPopularVideo 热门视频
func AddPopularVideo(ctx context.Context, videoID int64, score float64, createTime time.Time) error {
	CleanUpOldVideos(ctx)
	// 将视频ID和创建时间添加到一个Hash中
	RedisClient.HSet(ctx, "video_create_time", fmt.Sprintf("%d", videoID), fmt.Sprintf("%d", createTime.Unix()))

	// 使用 ZINCRBY 命令来增加或更新视频的分数
	intCmd := RedisClient.ZIncrBy(ctx, "popular_videos", score, fmt.Sprintf("%d", videoID))
	if intCmd.Err() != nil {
		fmt.Printf("更新视频热度失败: %v\n", intCmd.Err())
	}

	return nil
}

// CleanUpOldVideos 清除热门榜上七天前创建的视频
func CleanUpOldVideos(ctx context.Context) {
	// 获取当前时间7天前的Unix时间戳
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Unix()

	// 获取所有视频的创建时间
	videoCreateTimeMap := RedisClient.HGetAll(ctx, "video_create_time").Val()

	for videoID, createTimeStr := range videoCreateTimeMap {
		createTime, _ := strconv.ParseInt(createTimeStr, 10, 64)

		// 如果视频的创建时间早于7天前，那么删除这个视频
		if createTime < sevenDaysAgo {
			RedisClient.ZRem(ctx, "popular_videos", videoID)
			RedisClient.HDel(ctx, "video_create_time", videoID)
		}
	}
}

// GetTop30Videos 获取前三十个热门视频
func GetTop30Videos(ctx context.Context) ([]uint, error) {
	// 使用 ZREVRANGE 命令获取榜单中分数最高的前30个视频
	cmd := RedisClient.ZRevRangeWithScores(ctx, "popular_videos", 0, 29)
	if cmd.Err() != nil {
		fmt.Printf("获取热门视频失败: %v\n", cmd.Err())
		return nil, cmd.Err()
	}

	// 解析结果，获取视频ID
	videoIDs := make([]uint, len(cmd.Val()))
	for i, member := range cmd.Val() {
		videoID, err := strconv.ParseInt(member.Member.(string), 10, 64)
		if err != nil {
			fmt.Printf("解析视频ID失败: %v\n", err)
			return nil, err
		}
		videoIDs[i] = uint(videoID)
	}

	return videoIDs, nil
}

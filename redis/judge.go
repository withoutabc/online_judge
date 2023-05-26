package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"online_judge/util"
	"strconv"
)

func SetSubmissionId(ctx context.Context, submissionId int64) {
	// 设置键值对
	err := RDB.Set(ctx, strconv.FormatInt(submissionId, 10), "false", 0).Err()
	if err != nil {
		util.Log(err)
	}
}

func DeleteSubmissionId(ctx context.Context, submissionId int64) {
	// 删除键
	err := RDB.Del(ctx, strconv.FormatInt(submissionId, 10)).Err()
	if err != nil {
		util.Log(err)
	}
}

func GetValueAndExistence(ctx context.Context, submissionId int64) (string, bool) {
	// 获取指定key的值
	value, err := RDB.Get(ctx, strconv.FormatInt(submissionId, 10)).Result()
	if err == redis.Nil {
		// 指定key不存在
		return "", false
	} else if err != nil {
		// 获取值出错
		util.Log(err)
	}
	// 指定key存在，返回对应的值
	return value, true
}

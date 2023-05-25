package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func SetSubmissionId(ctx context.Context, submissionId int64) error {
	// 设置键值对
	err := RDB.Set(ctx, strconv.FormatInt(submissionId, 10), "false", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetValueAndExistence(ctx context.Context, submissionId int64) (string, bool, error) {
	// 获取指定key的值
	value, err := RDB.Get(ctx, strconv.FormatInt(submissionId, 10)).Result()
	if err == redis.Nil {
		// 指定key不存在
		return "", false, nil
	} else if err != nil {
		// 获取值出错
		return "", false, err
	}
	// 指定key存在，返回对应的值
	return value, true, nil
}

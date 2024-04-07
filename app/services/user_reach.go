package services

import (
	"context"
	"encoding/json"
	"gin-api/app/common/request"
	"gin-api/app/models"
	"gin-api/global"
	"strconv"
	"time"
)

const UserReachRedisQueueKey = "user_reach_queue"

type UserReachServiceImpl struct{}

func UserReachService() *UserReachServiceImpl {
	return &UserReachServiceImpl{}
}

func (s *UserReachServiceImpl) IsExistRecord(appcode string, userId int64, eventId int64) bool {
	var result = global.App.DB.Where("appcode = ?", appcode).Where("user_id = ?", userId).
		Where("event_id = ?", eventId).
		Select("id").First(&models.UserReachMeta{})
	if result.RowsAffected != 0 {
		return true
	}
	return false
}

func (s *UserReachServiceImpl) IsExistInRedisBit(params request.UserReach) bool {
	existKey := s.GetRedisBitKey(params.Appcode, params.EventId, params.Date)
	ctx := context.Background()
	redis := global.App.Redis
	if res, _ := redis.GetBit(ctx, existKey, params.UserId).Result(); res == 1 {
		return true
	}
	return false
}

func (s *UserReachServiceImpl) GetRedisBitKey(appcode string, eventId int64, Date string) string {
	return appcode + ":user_reach:" + Date + ":e" + strconv.FormatInt(eventId, 10)
}

func (s *UserReachServiceImpl) PushMetaToQueue(item request.UserReach) (err error) {
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	ctx := context.Background()
	redis := global.App.Redis
	_, err = redis.RPush(ctx, UserReachRedisQueueKey, data).Result()
	if err != nil {
		return
	}
	// cache
	bitKey := s.GetRedisBitKey(item.Appcode, item.EventId, item.Date)
	exists, err := redis.Exists(ctx, bitKey).Result()
	if err != nil {
		return
	}
	redis.SetBit(ctx, bitKey, item.UserId, 1)
	if exists == 0 {
		redis.Expire(ctx, bitKey, 24*time.Hour)
	}
	return
}

func (s *UserReachServiceImpl) FlushQueueToMysql(size int) int64 {
	values, err := global.App.Redis.LRange(context.Background(), UserReachRedisQueueKey, 0, int64(size-1)).Result()
	if err != nil {
		panic(err)
	}
	if len(values) == 0 {
		return 0
	}

	var data []models.UserReachMeta
	for _, value := range values {
		var meta models.UserReachMeta
		err := json.Unmarshal([]byte(value), &meta)
		if err != nil {
			panic(err)
		}
		data = append(data, meta)
	}

	insertCount := global.App.DB.Create(&data).RowsAffected
	if insertCount > 0 {
		global.App.Redis.LTrim(context.Background(), UserReachRedisQueueKey, insertCount, -1)
	}
	return insertCount
}

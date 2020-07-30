package boot

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"go-project-initial/configs"
	"math"
	"time"
)

func init() {
	_ = Redis()
}

var redisClient *redis.Client

func Redis() *redis.Client {
	if redisClient == nil {
		addr := fmt.Sprintf("%s:%d", configs.Conf.Redis.Host, configs.Conf.Redis.Port)
		redisClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: configs.Conf.Redis.Password, // no password set
			DB:       configs.Conf.Redis.DB,       // use default DB
		})
		pong, err := redisClient.Ping().Result()
		log.Println("redis ping:", pong)
		if err != nil && err != redis.Nil {
			panic(err)
		}
	}
	return redisClient
}

var ModeEnvMap = map[string]string{
	"debug":   "local",
	"test":    "test",
	"release": "production",
}

func RedisKey(key string) string {
	name := configs.Conf.App.Name
	mode := configs.Conf.App.Mode
	env := ModeEnvMap[mode]
	return fmt.Sprintf("%s-%s-%s", name, env, key)
}

//当前时间纳秒值少3位 16位
//1582802356200054928
//1582802356 秒
//200054928 纳秒
//200054
func RedisScore() float64 {
	return float64(time.Now().UnixNano() / 1000)
}

func RedisScoreToTime(score float64) (t time.Time) {
	t = time.Unix(int64(score/math.Pow10(6)), 0)
	return
}

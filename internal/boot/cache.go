package boot

import (
	"github.com/allegro/bigcache"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-project-initial/configs"
	"time"
)

var cache *bigcache.BigCache
var BigCacheConfig bigcache.Config

func init() {
	BigCacheConfig = bigcache.DefaultConfig(configs.Conf.App.CacheExpire)
	BigCacheConfig.CleanWindow = configs.Conf.App.CacheCleanInterval
	if configs.Conf.App.Mode != gin.ReleaseMode {
		BigCacheConfig.LifeWindow = 1 * time.Second
		BigCacheConfig.CleanWindow = 1 * time.Second
	}
}

func Cache() *bigcache.BigCache {
	if cache == nil {
		var err error
		cache, err = bigcache.NewBigCache(BigCacheConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
	return cache
}

func NewCache(config bigcache.Config) (cache *bigcache.BigCache) {
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(err)
	}
	return
}

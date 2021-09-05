package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"home/app/captcha/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCaptchaRepo)

// Data .
type Data struct {
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	// 初始化redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		Password: c.Redis.Password,
		DB: 0,
	})

	return &Data{rdb: rdb}, cleanup, nil
}

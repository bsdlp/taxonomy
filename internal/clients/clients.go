package clients

import (
	"github.com/bsdlp/taxonomy/internal/configuration"
	"github.com/redis/go-redis/v9"
)

type Clients struct {
	Redis redis.Cmdable
}

func New(cfg configuration.Object) (*Clients, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return &Clients{
		Redis: r,
	}, nil
}

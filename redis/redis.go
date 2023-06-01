package redis

import (
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

func NewTracingHook(opts ...redisotel.Option) *redisotel.TracingHook {
	return redisotel.NewTracingHook()
}

package redis

import (
	redisotel "github.com/go-redis/redis/extra/redisotel/v8"
)

func NewTracingHook(opts ...redisotel.Option) *redisotel.TracingHook {
	return redisotel.NewTracingHook()
}

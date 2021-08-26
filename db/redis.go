package db

import (
	"api_template/config"
	"github.com/gofiber/storage/redis"
)

func NewRedis() *redis.Storage  {
		rd := config.Config.Redis

	// Initialize custom config
	store := redis.New(redis.Config{
		Host:     rd.Addr,
		Port:     rd.Port,
		Username: rd.User,
		Password: rd.Password,
		Database: rd.DB,
		Reset:    false,
	})
	return store
}

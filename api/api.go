package api

import (
	"api_template/config"
	"api_template/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Api *fiber.App
var Store *session.Store

func InitApp() {

	Api = fiber.New()

	// 中间件设置
	Api.Use(recover.New())

	sotrage := db.NewRedis()
	Store = session.New(session.Config{
		Storage: sotrage,
		KeyLookup: config.Config.KeyLookup,
	})
}

func InitRoute() {

}

func Start() {
	Api.Listen(config.Config.SiteAddress)
}

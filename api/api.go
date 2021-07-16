package api

import (
	"api_template/config"
	"api_template/db"
	"api_template/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	"github.com/utahta/go-cronowriter"
	"os"
	"time"
)

var Api *iris.Application

func InitApp() {

	Api = iris.New()

	// 中间件设置
	Api.Use(recover.New())

	session := sessions.New(sessions.Config{
		Expires: time.Hour * 2,
		Cookie:  config.Config.SessionName,
	})
	session.UseDatabase(db.NewRedis())
	Api.Use(session.Handler())
	Api.Use(middleware.Cors)
}

func SetApplicationLog(path string) {
	w1 := cronowriter.MustNew(path +"/iris.log.%Y%m%d")
	Api.Logger().SetOutput(w1)
	Api.Logger().AddOutput(os.Stdout)
}

func InitRoute() {

}

func Start() {
	Api.Listen(config.Config.SiteAddress)
}

package main

import (
	"api_template/api"
	"api_template/config"
	"api_template/db"
	"api_template/utils"
	"fmt"
)

func main(){
	fmt.Println("读取配置文件")

	config.GetConfig("config/config.yaml")

	fmt.Println("初始化日志")

	utils.InitLogger(config.Config.LogPath)

	api.InitApp()

	api.InitRoute()

	db.Connect()

	db.AutoMigrate()

	api.Start()
}
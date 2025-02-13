package main

import (
	"bot/app"
	"bot/config"
	"bot/utils"
	"log"
)

func main() {
	// 加载配置
	configs := config.LoadConfig()
	// 全局日志格式定义
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 注册数据库ORM应用 - Ent
	//  clean() => defer client.Close()
	_, err := utils.InitEntClient(configs)
	if err != nil {
		log.Println("数据库初始化失败。", err)
	}
	// 注册Redis
	utils.InitRedis(configs)
	// 注册SMS应用
	utils.InitTencentSMS(configs)
	// 注册NowApi
	utils.InitNowApi(configs)
	// 注册定时任务服务端
	srv := utils.InitAsynqServer(configs)
	mux := app.InitTaskMux()
	if err := srv.Run(mux); err != nil {
		log.Println("Asynq 服务端启动失败：", err)
	}
}

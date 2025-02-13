package main

import (
	"bot/app"
	"bot/config"
	"bot/utils"
	"fmt"
	"log"
)

func main() {
	// 加载配置
	configs := config.LoadConfig()
	// 全局日志格式定义
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 注册Redis
	utils.InitRedis(configs)
	// 注册定时任务客户端
	utils.InitAsynqClient(configs)
	// 注册所有定时监控任务
	app.InitTaskMonitor()
	if err := utils.AsynqScheduler.Run(); err != nil {
		fmt.Println("Asynq scheduler 客户端启动失败。")
	}
}

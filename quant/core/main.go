package main

import (
	"core/app"
	"core/config"
	"core/utils"
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
		log.Println("数据库初始化失败：", err)
	}
	// 注册Redis
	utils.InitRedis(configs)
	// 注册Web应用 - Gin
	r := utils.InitGin(configs)
	// 注册CORS应用 - gin-cors
	utils.InitGinCORS(r, configs)
	// 注册JWT应用 - gin-jwt
	authMiddleware := utils.InitGinJwt(configs)
	// 注册路由
	app.InitRoute(r, authMiddleware)
	err = r.Run(":8081")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

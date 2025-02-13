package utils

import (
	"core/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinCORS(r *gin.Engine, c config.Configs) {
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     c.CORS.AllowOrigins,
		AllowMethods:     c.CORS.AllowMethods,
		AllowHeaders:     c.CORS.AllowHeaders,
		ExposeHeaders:    c.CORS.ExposeHeaders,
		AllowCredentials: c.CORS.AllowCredentials,
		MaxAge:           c.CORS.MaxAge,
	})
	r.Use(corsMiddleware)
}

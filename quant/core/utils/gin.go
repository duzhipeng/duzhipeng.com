package utils

// Web框架：Gin
// https://gin-gonic.com

import (
	"core/config"
	"github.com/gin-gonic/gin"
)

func InitGin(c config.Configs) *gin.Engine {
	debug := c.Debug
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	return r
}

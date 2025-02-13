package app

import (
	"core/app/api/account"
	"core/app/api/job"
	"core/app/upgrade"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	// Upgrade
	Upgrade := r.Group("/upgrade")
	Upgrade.Use(authMiddleware.MiddlewareFunc())
	{
		Upgrade.POST("", upgrade.UpToHandler)
	}
	// Account
	accountGroup := r.Group("/account")
	accountGroup.Use(authMiddleware.MiddlewareFunc())
	{
		accountGroup.POST("", account.PostAccountRecordHandler)
	}
	// Job
	jobGroup := r.Group("/job")
	jobGroup.GET("/check_order_station", job.GetCheckOrderStationRecordHandler)
	jobGroup.Use(authMiddleware.MiddlewareFunc())
	{

	}

}

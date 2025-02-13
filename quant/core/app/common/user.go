package common

import (
	"core/ent"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetLoginUserInfo 从登录会话取用户信息
func GetLoginUserInfo(ctx *gin.Context, client *ent.Client) (userData *ent.User, error error) {
	currentUser, exists := ctx.Get("currentUser")
	loginUser, ok := currentUser.(*ent.User)
	if !exists || !ok {
		return nil, fmt.Errorf("请登录。")
	}
	return loginUser, nil
}

package utils

import (
	"core/config"
	"core/ent"
	"core/ent/user"
	"core/locale"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var identityKey string

func InitGinJwt(c config.Configs) *jwt.GinJWTMiddleware {
	identityKey = c.GinJwt.IdentityKey
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           c.GinJwt.Realm,      // 可以理解成该中间件的名称，用于展示，默认值为 gin jwt
		Key:             c.GinJwt.Key,        // 服务端密钥
		Timeout:         c.GinJwt.Timeout,    // Token过期时间，默认值为time.Hour
		MaxRefresh:      c.GinJwt.MaxRefresh, // Token 更新时间
		IdentityKey:     identityKey,         // 身份验证的Key值，默认值为identity
		PayloadFunc:     GinJWTPayloadFunc,
		Authenticator:   GinJWTAuthenticator,
		Authorizator:    GinJWTAuthorizator,
		Unauthorized:    GinJWTUnauthorized,
		LoginResponse:   GinJWTResponse,
		RefreshResponse: GinJWTResponse,
		// COOKIE设置
		SendCookie:     true,
		CookieHTTPOnly: true,                  // JS can't modify
		SecureCookie:   c.GinJwt.SecureCookie, //non HTTPS dev environments
		CookieDomain:   c.GinJwt.CookieDomain,
		CookieName:     c.GinJwt.CookieName, // default jwt
		CookieSameSite: c.GinJwt.CookieSameSite,
		TokenLookup:    c.GinJwt.TokenLookup,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware
}

type PostLoginField struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// GinJWTPayloadFunc 添加额外业务相关的信息
func GinJWTPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*ent.User); ok {
		return jwt.MapClaims{
			identityKey: v.ID,
		}
	}
	return jwt.MapClaims{}
}

// GinJWTAuthenticator 根据登录信息对用户进行身份验证的回调函数
func GinJWTAuthenticator(c *gin.Context) (interface{}, error) {
	var loginField PostLoginField
	if err := c.ShouldBindJSON(&loginField); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginField.Username
	password := loginField.Password
	// 检查用户名
	u, err := DBClient.User.Query().Where(user.Username(username)).Only(c)
	if err != nil {
		return nil, fmt.Errorf(locale.Result.Unauthorized.Message)
	}
	// 检查用户名密码是否匹配
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return nil, fmt.Errorf(locale.Result.Unauthorized.Message)
	} else {
		return u, nil
	}
}

// GinJWTAuthorizator 登录后其他接口验证传入的 token 方法，可以用于角色权限控制
func GinJWTAuthorizator(data interface{}, c *gin.Context) bool {
	id, ok := data.(float64)
	if ok {
		u, err := DBClient.User.Query().Where(user.ID(int(id))).Only(c)
		if err != nil {
			return false
		}
		// 当前会话的用户标识
		c.Set("currentUser", u)
		// 检查用户状态
		if u.Status != "BANNED" {
			return true
		}
	}
	return false
}

// GinJWTUnauthorized 验证失败时系统调用此函数一般返回错误信息
func GinJWTUnauthorized(c *gin.Context, code int, message string) {
	res := locale.Result.Unauthorized
	res.Message = "用户名或密码错误。"
	c.JSON(code, res)
}

// GinJWTResponse 定制 Response
func GinJWTResponse(c *gin.Context, code int, token string, t time.Time) {
	var res locale.ResponseField
	data := gin.H{
		"token":  token,
		"expire": t,
	}
	// 正常返回
	res = locale.Result.Ok
	res.Message = "欢迎回来。"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

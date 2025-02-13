package config

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

type Configs struct {
	// Debug MODE:Ture is debug.
	Debug bool
	// DB PostgreSQL
	DB struct {
		Type   string
		Source string
	}
	// Redis
	Redis struct {
		URI      string
		Password string
		DB       int
	}
	// Official CORS gin's middleware https://github.com/gin-contrib/cors
	CORS struct {
		AllowOrigins     []string
		AllowMethods     []string
		AllowHeaders     []string
		ExposeHeaders    []string
		AllowCredentials bool
		MaxAge           time.Duration
	}
	// gin's middleware GinJwt https://github.com/appleboy/gin-jwt
	GinJwt struct {
		Key            []byte
		Realm          string
		Timeout        time.Duration
		MaxRefresh     time.Duration
		IdentityKey    string
		TokenHeadName  string
		SecureCookie   bool
		CookieDomain   string
		CookieSameSite http.SameSite
		CookieName     string
		TokenLookup    string
	}
	// Tencent cloud SMS https://cloud.tencent.com/document/product/382/43199
	SMS struct {
		SecretID  string
		SecretKey string
	}
	// AWS S3
	AWSS3 struct {
		AccessKeyID     string
		SecretAccessKey string
		Endpoint        string
		Region          string
	}
	// Alipay
	Alipay struct {
		AppID        string
		PrivateKey   string
		IsProduction bool
	}
}

func DevConfig(debug bool) Configs {
	var configs Configs
	configs.Debug = debug
	configs.DB.Type = os.Getenv("DB_TYPE")
	configs.DB.Source = os.Getenv("DB_SOURCE")
	configs.Redis.URI = os.Getenv("REDIS_URI")
	configs.Redis.Password = os.Getenv("REDIS_PWD")
	configs.Redis.DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	configs.CORS.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5050"}
	configs.CORS.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	configs.CORS.AllowHeaders = []string{"Content-Type", "X-CSRF-TOKEN", "Authorization", "Access-Control-Allow-Headers"}
	configs.CORS.ExposeHeaders = []string{"Content-Length"}
	configs.CORS.AllowCredentials = true
	configs.CORS.MaxAge = 24 * time.Hour
	configs.GinJwt.Key = []byte(os.Getenv("SECRET_KEY"))
	configs.GinJwt.Realm = "jwt"
	configs.GinJwt.Timeout = time.Hour * 8
	configs.GinJwt.MaxRefresh = time.Hour * 24
	configs.GinJwt.IdentityKey = "token"
	configs.GinJwt.TokenLookup = "header: Authorization, query: token, cookie: token"
	configs.GinJwt.CookieName = "token"
	configs.GinJwt.SecureCookie = !debug
	configs.GinJwt.CookieDomain = "localhost"
	configs.GinJwt.CookieSameSite = http.SameSiteLaxMode
	configs.SMS.SecretID = os.Getenv("TENCENTCLOUD_SECRET_ID")
	configs.SMS.SecretKey = os.Getenv("TENCENTCLOUD_SECRET_KEY")
	configs.AWSS3.AccessKeyID = os.Getenv("AWS_S3_ACCESS_KEY_ID")
	configs.AWSS3.SecretAccessKey = os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	configs.AWSS3.Endpoint = os.Getenv("AWS_S3_ENDPOINT")
	configs.AWSS3.Region = os.Getenv("AWS_S3_REGION")
	configs.Alipay.AppID = os.Getenv("ALIPAY_APPID")
	configs.Alipay.PrivateKey = os.Getenv("ALIPAY_PKEY")
	configs.Alipay.IsProduction = true
	return configs
}

func LiveConfig(debug bool) Configs {
	var configs Configs
	configs.Debug = debug
	configs.DB.Type = os.Getenv("DB_TYPE")
	configs.DB.Source = os.Getenv("DB_SOURCE")
	configs.Redis.URI = os.Getenv("REDIS_URI")
	configs.Redis.Password = os.Getenv("REDIS_PWD")
	configs.Redis.DB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	configs.CORS.AllowOrigins = []string{"https://www.qeg.net", "https://quant.qeg.net", "https://www.qeg.net.cn"}
	configs.CORS.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	configs.CORS.AllowHeaders = []string{"Content-Type", "X-CSRF-TOKEN", "Authorization", "Access-Control-Allow-Headers"}
	configs.CORS.ExposeHeaders = []string{"Content-Length"}
	configs.CORS.AllowCredentials = true
	configs.CORS.MaxAge = 24 * time.Hour
	configs.GinJwt.Key = []byte(os.Getenv("SECRET_KEY"))
	configs.GinJwt.Realm = "jwt"
	configs.GinJwt.Timeout = time.Hour * 8
	configs.GinJwt.MaxRefresh = time.Hour * 24
	configs.GinJwt.IdentityKey = "token"
	configs.GinJwt.TokenLookup = "header: Authorization, query: token, cookie: token"
	configs.GinJwt.CookieName = "token"
	configs.GinJwt.SecureCookie = !debug
	configs.GinJwt.CookieDomain = ".qeg.net"
	configs.GinJwt.CookieSameSite = http.SameSiteLaxMode
	configs.SMS.SecretID = os.Getenv("TENCENTCLOUD_SECRET_ID")
	configs.SMS.SecretKey = os.Getenv("TENCENTCLOUD_SECRET_KEY")
	configs.AWSS3.AccessKeyID = os.Getenv("AWS_S3_ACCESS_KEY_ID")
	configs.AWSS3.SecretAccessKey = os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	configs.AWSS3.Endpoint = os.Getenv("AWS_S3_ENDPOINT")
	configs.AWSS3.Region = os.Getenv("AWS_S3_REGION")
	configs.Alipay.AppID = os.Getenv("ALIPAY_APPID")
	configs.Alipay.PrivateKey = os.Getenv("ALIPAY_PKEY")
	configs.Alipay.IsProduction = true
	return configs
}

func LoadConfig() Configs {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	env := os.Getenv("MODE")
	switch env {
	case "DEV":
		return DevConfig(debug)
	case "LIVE":
		return LiveConfig(debug)
	default:
		return DevConfig(debug)
	}

}

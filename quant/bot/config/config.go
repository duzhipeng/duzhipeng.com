package config

import (
	"os"
	"strconv"
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
	// Tencent cloud SMS https://cloud.tencent.com/document/product/382/43199
	SMS struct {
		SecretID  string
		SecretKey string
	}
	// NowAPi
	NowApi struct {
		Key  string
		Sign string
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
	configs.SMS.SecretID = os.Getenv("TENCENTCLOUD_SECRET_ID")
	configs.SMS.SecretKey = os.Getenv("TENCENTCLOUD_SECRET_KEY")
	configs.NowApi.Key = os.Getenv("NOWAPI_KEY")
	configs.NowApi.Sign = os.Getenv("NOWAPI_SIGN")

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
	configs.SMS.SecretID = os.Getenv("TENCENTCLOUD_SECRET_ID")
	configs.SMS.SecretKey = os.Getenv("TENCENTCLOUD_SECRET_KEY")
	configs.NowApi.Key = os.Getenv("NOWAPI_KEY")
	configs.NowApi.Sign = os.Getenv("NOWAPI_SIGN")

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

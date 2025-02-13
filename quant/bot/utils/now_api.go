package utils

import "bot/config"

var NowApi APIKeyAndSign

type APIKeyAndSign struct {
	URL    string
	AppKey string
	Sign   string
}

func InitNowApi(c config.Configs) {
	NowApi.URL = "https://sapi.k780.com"
	NowApi.AppKey = c.NowApi.Key
	NowApi.Sign = c.NowApi.Sign
}

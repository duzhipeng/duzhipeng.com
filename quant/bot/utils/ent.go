package utils

// 数据库ORM：Ent
// https://entgo.io

import (
	"bot/config"
	"bot/ent"
	_ "github.com/lib/pq"
)

var DBClient *ent.Client

func InitEntClient(c config.Configs) (func(), error) {
	client, err := ent.Open(c.DB.Type, c.DB.Source)
	if err != nil {
		return nil, err
	}
	// defer client.Close()
	clean := func() {
		err := client.Close()
		if err != nil {
			return
		}
	}
	DBClient = client
	return clean, err
}

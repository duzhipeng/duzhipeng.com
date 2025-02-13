package utils

// 数据库ORM：Ent
// https://entgo.io

import (
	"context"
	"core/config"
	"core/ent"
	"core/ent/migrate"
	_ "github.com/lib/pq"
)

var DBClient *ent.Client

func InitEntClient(c config.Configs) (func(), error) {
	client, err := ent.Open(c.DB.Type, c.DB.Source)
	if err != nil {
		return nil, err
	}
	// Run migration.
	ctx := context.Background()
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
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

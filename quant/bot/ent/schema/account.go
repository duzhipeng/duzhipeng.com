package schema

// 更新数据结构后，运行：go generate ./ent
import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Annotations of the Account.
func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account"},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(32),          // 用户名
		field.String("password").MaxLen(32),          // 密码
		field.String("token").MaxLen(128).Optional(), // token
		// 通用部份
		field.Time("created_at").Default(time.Now).Immutable(),             // 创建时间
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), // 更新时间
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{}
}

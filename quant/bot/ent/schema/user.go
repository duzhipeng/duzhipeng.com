package schema

// 更新数据结构后，运行：go generate ./ent
import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(16).Unique(),         // 用户名
		field.String("phone").MaxLen(16).Optional().Unique(), // 电话号码
		field.String("email").MaxLen(128).Optional(),         // 邮箱地址
		field.String("password_hash").MaxLen(128).Sensitive(),
		field.Enum("status").
			Values("PENDING", "ACTIVE", "BANNED").Default("PENDING"), // 状态：待激活（默认），已激活（会员），封禁
		// 通用部份
		field.Time("created_at").Default(time.Now).Immutable(),             // 创建时间
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), // 更新时间
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

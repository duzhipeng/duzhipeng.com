package schema

// 更新数据结构后，运行：go generate ./ent
import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Station holds the schema definition for the Station entity.
type Station struct {
	ent.Schema
}

// Annotations of the Station.
func (Station) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "station"},
	}
}

// Fields of the Station.
func (Station) Fields() []ent.Field {
	return []ent.Field{
		field.Int("agencyId").Unique(),
		field.String("agencyName").MaxLen(256),
		// 通用部份
		field.Time("created_at").Default(time.Now).Immutable(),             // 创建时间
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), // 更新时间
	}
}

// Edges of the Station.
func (Station) Edges() []ent.Edge {
	return []ent.Edge{}
}

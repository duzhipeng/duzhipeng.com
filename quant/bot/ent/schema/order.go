package schema

// 更新数据结构后，运行：go generate ./ent
import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Annotations of the Order.
func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order"},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("maintOrderNo").Unique().MaxLen(64),
		field.String("vehiclePlateNo").MaxLen(64),
		field.String("vehicleTeamName").MaxLen(64),
		field.String("maintRequestType").MaxLen(64),
		field.Int("dispatchedStationId"),
		field.String("stationName").MaxLen(128),
		// 通用部份
		field.Time("created_at").Default(time.Now).Immutable(),             // 创建时间
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now), // 更新时间
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{}
}

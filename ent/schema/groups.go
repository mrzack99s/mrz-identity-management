package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Groups holds the schema definition for the Groups entity.
type Groups struct {
	ent.Schema
}

// Fields of the Groups.
func (Groups) Fields() []ent.Field {
	return []ent.Field{
		field.String("g_name").NotEmpty(),
		field.Bool("g_is_int_org").Default(true),
		field.Bool("g_is_super_admin").Default(false),
		field.Time("g_created_at").Default(time.Now).Optional(),
	}
}

// Edges of the Groups.
func (Groups) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("use_bandwidth", GroupBandwidth.Type).
			Ref("groups").
			Unique(),
		edge.To("users", Users.Type),
	}
}

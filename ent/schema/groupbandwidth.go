package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupBandwidth holds the schema definition for the GroupBandwidth entity.
type GroupBandwidth struct {
	ent.Schema
}

// Fields of the GroupBandwidth.
func (GroupBandwidth) Fields() []ent.Field {
	return []ent.Field{
		field.Int("gbw_download_speed"),
		field.Int("gbw_upload_speed"),
		field.Time("gbw_created_at").Default(time.Now).Optional(),
	}
}

// Edges of the GroupBandwidth.
func (GroupBandwidth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Groups.Type),
	}
}

package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OnlineSession holds the schema definition for the OnlineSession entity.
type OnlineSession struct {
	ent.Schema
}

// Fields of the OnlineSession.
func (OnlineSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip_address").MaxLen(15).NotEmpty().NotEmpty().Match(regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)),
		field.String("u_pid").MaxLen(13).NotEmpty().Match(regexp.MustCompile("[a-zA-Z0-9]+")).Unique(),
	}
}

// Edges of the OnlineSession.
func (OnlineSession) Edges() []ent.Edge {
	return nil
}

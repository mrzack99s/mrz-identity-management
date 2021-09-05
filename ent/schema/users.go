package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("u_pid").MaxLen(13).NotEmpty().Match(regexp.MustCompile("[a-zA-Z0-9]+")).Unique(),
		field.String("u_orgid").MaxLen(20).NotEmpty().Match(regexp.MustCompile("[a-zA-Z0-9]+")).Unique(),
		field.String("u_first_name").MaxLen(50).NotEmpty().Match(regexp.MustCompile("[a-zA-Zก-๙]*")),
		field.String("u_last_name").MaxLen(50).NotEmpty().Match(regexp.MustCompile("[a-zA-Zก-๙]*")),
		field.Bool("u_is_active").Default(false),
		field.Time("u_created_at").Default(time.Now).Immutable().Optional(),
		field.Time("u_password_updated_at").Optional(),
		field.Time("u_expired_at").Optional(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("in_group", Groups.Type).
			Ref("users").
			Unique(),
	}
}

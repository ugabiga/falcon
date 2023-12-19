package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Authentication holds the schema definition for the Authentication entity.
type Authentication struct {
	ent.Schema
}

// Fields of the Authentication.
func (Authentication) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.Enum("provider").
			Values("google", "facebook"),
		field.String("identifier").
			Unique(),
		field.String("credential").
			Sensitive(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Authentication.
func (Authentication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("authentications").
			Unique(),
	}
}

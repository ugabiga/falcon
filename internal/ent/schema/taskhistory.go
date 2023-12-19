package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TaskHistory holds the schema definition for the TaskHistory entity.
type TaskHistory struct {
	ent.Schema
}

// Fields of the TaskHistory.
func (TaskHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive(),
		field.Bool("is_success"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the TaskHistory.
func (TaskHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("task_histories").
			Unique(),
	}
}

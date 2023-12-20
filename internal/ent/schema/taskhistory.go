package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Int("id").
			Positive().
			GoType(int(0)),
		field.Int("task_id").
			Positive().
			GoType(int(0)),
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
			Unique().
			Required().
			Field("task_id"),
	}
}

func (TaskHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			GoType(int(0)),
		field.Int("trading_account_id").
			Positive().
			GoType(int(0)),
		field.String("currency"),
		field.Float("size").
			Default(0).
			Comment("size of crypto currency to buy/sell"),
		field.String("symbol").
			Comment("symbol of currency to buy/sell"),
		field.String("cron"),
		field.Time("next_execution_time").
			Optional(),
		field.Bool("is_active").
			Default(true),
		field.String("type"),
		field.JSON("params", map[string]interface{}{}).
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("trading_account", TradingAccount.Type).
			Ref("tasks").
			Unique().
			Field("trading_account_id").
			Required(),
		edge.To("task_histories", TaskHistory.Type),
	}
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

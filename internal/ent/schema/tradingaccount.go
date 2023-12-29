package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TradingAccount holds the schema definition for the TradingAccount entity.
type TradingAccount struct {
	ent.Schema
}

// Fields of the TradingAccount.
func (TradingAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			GoType(int(0)),
		field.Int("user_id").
			Positive().
			GoType(int(0)),
		field.String("name"),
		field.String("exchange"),
		field.String("ip"),
		field.String("key").
			Unique(),
		field.String("secret").
			Sensitive(),
		field.String("phrase").
			Sensitive().
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the TradingAccount.
func (TradingAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("trading_accounts").
			Unique().
			Field("user_id").
			Required(),
		edge.To("tasks", Task.Type),
	}
}
func (TradingAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

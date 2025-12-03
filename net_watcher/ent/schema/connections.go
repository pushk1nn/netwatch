package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Connections holds the schema definition for the Connections entity.
type Connections struct {
	ent.Schema
}

// Fields of the Connections.
func (Connections) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("Unique ID of individual event").
			Default(uuid.New),
		field.Time("time").
			Comment("Time at which event occurred"),
		field.Int64("unix_time").
			Comment("Time at which event occurred as Unix Time Stamp"),
		field.String("type").
			Comment("Event type (connect, disconnect)"),
		field.String("ip").
			Comment("IP from which a connection is coming from"),
	}
}

// Edges of the Connections.
func (Connections) Edges() []ent.Edge {
	return nil
}

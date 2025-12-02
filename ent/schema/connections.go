package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Connections holds the schema definition for the Connections entity.
type Connections struct {
	ent.Schema
}

// Fields of the Connections.
func (Connections) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").
			Comment("Unique ID associated with individual log").
			Unique(),
		field.Time("start_time").
			Comment("Time at which connection started"),
		field.Time("end_time").
			Comment("Time at which connection ended"),
		field.String("ip").
			Comment("IP from which a connection is coming from"),
	}
}

// Edges of the Connections.
func (Connections) Edges() []ent.Edge {
	return nil
}

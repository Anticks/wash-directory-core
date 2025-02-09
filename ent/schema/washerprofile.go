package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type WasherProfile struct {
	ent.Schema
}

func (WasherProfile) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}).
		Default(uuid.New).
		Unique().
		Immutable(),

		field.String("service_details").Optional(),
		field.String("availability").Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (WasherProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("washer_profile").
			Unique().
			Required(),
	}
}

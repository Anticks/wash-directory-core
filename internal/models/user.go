package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// UUID primary key
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New). // generate new UUID on creation
			Unique().
			Immutable(),

		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{

		// 1-to-1 with WasherProfile
		edge.To("washer_profile", WasherProfile.Type).
			Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// User holds the schema definition for the User entity.
type Link struct {
	ent.Schema
}

// Fields of the User.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("url").NotEmpty(),
		field.String("hash").NotEmpty(),
		field.Text("describe").Optional(),
		field.JSON("json", domain.Link{}),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the User.
func (Link) Edges() []ent.Edge {
	return nil
}

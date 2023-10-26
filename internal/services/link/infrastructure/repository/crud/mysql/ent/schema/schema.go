package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
)

// User holds the schema definition for the User entity.
type Link struct {
	ent.Schema
}

// Annotations of the Link.
func (Link) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "links"},
		entsql.WithComments(true),
		schema.Comment("Link holds the schema definition for the Link entity."),
	}
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUID"),
		field.String("url").
			Comment("URL").
			NotEmpty(),
		field.String("hash").
			Comment("Hash").
			NotEmpty(),
		field.Text("describe").
			Comment("Describe").
			Optional(),
		field.JSON("json", domain.Link{}).
			Comment("JSON"),
		field.Time("created_at").
			Comment("Created at"),
		field.Time("updated_at").
			Comment("Updated at"),
	}
}

// Indexes of the Link.
func (Link) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("url", "url").Unique(),
		index.Fields("hash", "hash").Unique(),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return nil
}

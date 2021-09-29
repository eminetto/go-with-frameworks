package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Lesson holds the schema definition for the Lesson entity.
type Lesson struct {
	ent.Schema
}

// Fields of the Lesson.
func (Lesson) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id"),
		field.String("name"),
	}
}

// Edges of the Lesson.
func (Lesson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("lessons").
			Unique(),
	}
}

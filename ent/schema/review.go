package schema

import "entgo.io/ent"

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
        field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
				field.UUID("user_uuid"),
				field.UUID("tool_uuid"),
        field.Int("rating"),
				field.String("comment"),
				field.Time("created_at").Default(time.Now),
    }
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return nil
}

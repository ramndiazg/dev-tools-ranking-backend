package schema

import "entgo.io/ent"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
        field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
        field.String("name").Default("unknown"),
				field.String("username").Unique(),
				field.String("email").Unique(),
				field.String("password_hash"),
				field.Time("created_at").Default(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

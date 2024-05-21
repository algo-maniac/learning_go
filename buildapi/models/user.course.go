package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user in the MongoDB database.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}


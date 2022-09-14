package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// -----------------------------------------------
// -- User models & dtos
// -----------------------------------------------

type CreateUserDto struct {
	Name string `json:"name" binding:"required"`
}

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `json:"name"`
}

type Users []User

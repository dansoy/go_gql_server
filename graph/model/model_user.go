package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//NewUser struct
type NewUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//User struct
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
}

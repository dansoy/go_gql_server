package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//NewTodo struct
type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"user_id"`
}

//Todo struct
type Todo struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Text   string             `json:"text"`
	Done   bool               `json:"done"`
	UserID string             `json:"user"`
}

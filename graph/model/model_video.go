package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//NewVideo struct
type NewVideo struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	UserID string `json:"user_id"`
}

//Video struct
type Video struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Title  string             `json:"title"`
	URL    string             `json:"url"`
	UserID string             `json:"user"`
}

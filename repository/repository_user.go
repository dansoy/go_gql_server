package repository

import (
	"context"
	"log"

	"github.com/dansoy/go_gql_server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserRepository interface
type UserRepository interface {
	Save(user *model.User)
	Find() []*model.User
	FindByID(id string) *model.User
}

func (db *databaseUser) Save(user *model.User) {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *databaseUser) Find() []*model.User {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.User
	for cursor.Next(context.TODO()) {
		var v *model.User
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}

func (db *databaseUser) FindByID(id string) *model.User {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	var result *model.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

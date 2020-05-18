package repository

import (
	"context"
	"log"

	"github.com/dansoy/go_gql_server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TodoRepository interface
type TodoRepository interface {
	Save(todo *model.Todo)
	Find() []*model.Todo
	FindByID(id string) *model.Todo
}

func (db *databaseTodo) Save(todo *model.Todo) {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	_, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *databaseTodo) Find() []*model.Todo {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Todo
	for cursor.Next(context.TODO()) {
		var v *model.Todo
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}

func (db *databaseTodo) FindByID(id string) *model.Todo {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	var result *model.Todo
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

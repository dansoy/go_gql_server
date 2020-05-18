package repository

import (
	"context"
	"log"

	"github.com/dansoy/go_gql_server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//VideoRepository interface
type VideoRepository interface {
	Save(video *model.Video)
	Find() []*model.Video
	FindByID(id string) *model.Video
}

func (db *databaseVideo) Save(video *model.Video) {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	_, err := collection.InsertOne(context.TODO(), video)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *databaseVideo) Find() []*model.Video {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.Video
	for cursor.Next(context.TODO()) {
		var v *model.Video
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}

func (db *databaseVideo) FindByID(id string) *model.Video {
	collection := db.client.Database(db.dbName).Collection(db.cName)
	var result *model.Video
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

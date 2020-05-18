package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type databaseVideo struct {
	client *mongo.Client
	dbName string
	cName  string
}

type databaseTodo struct {
	client *mongo.Client
	dbName string
	cName  string
}

type databaseUser struct {
	client *mongo.Client
	dbName string
	cName  string
}

//
const (
	DATABASE         = "gqlserver"
	COLLECTIONVIDEOS = "videos"
	COLLECTIONTODOS  = "todos"
	COLLECTIONUSERS  = "users"
)

//
var (
	database  map[string]interface{}
	dbclient  *mongo.Client
	TodoRepo  TodoRepository
	VideoRepo VideoRepository
	UserRepo  UserRepository
)

func init() {
	dbclient = connectRepository()
	TodoRepo = GetTodoRepository()
	VideoRepo = GetVideoRepository()
	UserRepo = GetUserRepository()
}

func connectRepository() *mongo.Client {
	// mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	MONGODB := os.Getenv("MONGOGQLSERVER")

	// Set client options
	clientOptions := options.Client().ApplyURI(MONGODB)

	clientOptions = clientOptions.SetMaxPoolSize(50)

	// Set timeout
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Connect to MongoDB
	userClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = userClient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return userClient
}

//GetTodoRepository function, returns TodoRepository
func GetTodoRepository() TodoRepository {
	return &databaseTodo{
		client: dbclient,
		dbName: DATABASE,
		cName:  COLLECTIONTODOS,
	}
}

//GetVideoRepository function, returns VideoRepository
func GetVideoRepository() VideoRepository {
	return &databaseVideo{
		client: dbclient,
		dbName: DATABASE,
		cName:  COLLECTIONVIDEOS,
	}
}

//GetUserRepository function, returns UserRepository
func GetUserRepository() UserRepository {
	return &databaseUser{
		client: dbclient,
		dbName: DATABASE,
		cName:  COLLECTIONUSERS,
	}
}

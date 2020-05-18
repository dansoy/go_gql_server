# go_gql_server
Go Language and MongoDB GraphQL Server

# Project Usage
	-Create folder
		$ mkdir PROJECT_NAME
		$ cd PROJECT_NAME
	- Initialize Go Module
		$ go mod init github.com/[username]/PROJECT_NAME
	- Install dependency gqlgen and mongo-go-driver
		$ go get github.com/99designs/gqlgen
		$ go get go.mongodb.org/mongo-driver/mongo
	- Create the project skeleton
		$ go run github.com/99designs/gqlgen init
	- Modify gqlgen.yml if necessary
	- Define schema in graph/schema.graphql
	- Generate model and schema resolver (Run this every time you change the GraphQL Schema)
		$ gqlgen generate
		or
		$ go generate ./...
	- Run server
		$ go run server.go

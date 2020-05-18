package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/dansoy/go_gql_server/graph/model"

// The file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver struct
type Resolver struct {
	todos  []*model.Todo
	videos []*model.Video
	users  []*model.User
}

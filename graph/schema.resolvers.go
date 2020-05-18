package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dansoy/go_gql_server/graph/generated"
	"github.com/dansoy/go_gql_server/graph/model"
	"github.com/dansoy/go_gql_server/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:     primitive.NewObjectID(),
		Text:   input.Text,
		UserID: input.UserID,
	}
	repository.TodoRepo.Save(todo)
	return todo, nil
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	video := &model.Video{
		ID:     primitive.NewObjectID(),
		Title:  input.Title,
		URL:    input.URL,
		UserID: input.UserID,
	}
	repository.VideoRepo.Save(video)
	return video, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:        primitive.NewObjectID(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	repository.UserRepo.Save(user)
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return repository.TodoRepo.Find(), nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return repository.VideoRepo.Find(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return repository.UserRepo.Find(), nil
}

func (r *subscriptionResolver) TodoAdded(ctx context.Context, repoFullName string) (<-chan *model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VideoAdded(ctx context.Context, repoFullName string) (<-chan *model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserAdded(ctx context.Context, repoFullName string) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return repository.UserRepo.FindByID(obj.UserID), nil
}

func (r *videoResolver) User(ctx context.Context, obj *model.Video) (*model.User, error) {
	return repository.UserRepo.FindByID(obj.UserID), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// Video returns generated.VideoResolver implementation.
func (r *Resolver) Video() generated.VideoResolver { return &videoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }

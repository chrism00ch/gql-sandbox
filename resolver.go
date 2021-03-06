package gql_sandbox

import (
	"context"
	"fmt"
	"github.com/chrism00ch/gql-sandbox/gettingstarted"
	"math/rand"
)

type Resolver struct {
	todos []gettingstarted.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (gettingstarted.Todo, error) {
	todo := gettingstarted.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]gettingstarted.Todo, error) {
	return r.todos, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *gettingstarted.Todo) (User, error) {
	return User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	modelgen "github.com/Chanra-IG/go-core/src/graphql/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input modelgen.NewTodo) (*modelgen.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*modelgen.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

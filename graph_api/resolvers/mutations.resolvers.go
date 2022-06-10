package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/gen"
	"github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/model"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, inputData model.OrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

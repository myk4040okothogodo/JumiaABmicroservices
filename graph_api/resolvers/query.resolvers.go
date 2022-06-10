package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/gen"
	"github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/model"
)

func (r *queryResolver) Orders(ctx context.Context, id *string, email *string, weight *string, phonenumber *string, countrycode *string) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

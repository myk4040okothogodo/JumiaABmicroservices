package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
  "log"
	"context"
  service_bv1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B"
	"github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/gen"
	"github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/model"
)

func (r *queryResolver) Orders(ctx context.Context, id *string, email *string, weight *string, phonenumber *string, countrycode *string) ([]*model.Order, error) {
	orders := []*model.Order{}

  if id != nil {
      log.Printf("Order id: %s", *id)
      GetOrderResponse, err := r.services.ServiceB().GetOrder(ctx, &service_bv1.GetOrderRequest{Id: *id})
      if err != nil {
          return nil, err
      }
      orders = append(orders, service2GraphOrder(GetOrderResponse.Order))
  
  } else if weight != nil {
      log.Printf("Order weight: %s", *weight)
      GetOrdersByWeightResponse, err := r.services.ServiceB().GetOrdersByWeight(ctx, &service_bv1.GetOrdersByWeightRequest{Weight: *weight})
      if err != nil {
          return nil, err
      }
      for _, order := range GetOrdersByWeightResponse.Orders{
          orders = append(orders, service2GraphOrder(order))
      }

  } else if countrycode != nil {
    log.Printf("Order countrycode: %s", *countrycode)
    GetCountryOrdersResponse, err := r.services.ServiceB().GetAllCountryOrders(ctx, &service_bv1.GetCountryOrdersRequest{Countrycode: *countrycode} )
    if err != nil {
        return nil, err
    }
    for _, order := range GetCountryOrdersResponse.Orders{
        orders = append(orders, service2GraphOrder(order))
    }

  } else if email != nil {
    log.Printf("Order email: %s", *email)
    GetOrderByEmailResponse, err := r.services.ServiceB().GetOrderByEmail(ctx, &service_bv1.GetOrderByEmailRequest{Email: *email})
    if err != nil {
        return nil, err
    }
    orders = append(orders, service2GraphOrder(GetOrderByEmailResponse.Order))
  
  } else {
    GetOrdersRequest, err := r.services.ServiceB(). GetAllOrders(ctx, &service_bv1.GetOrdersRequest{})
    if err != nil {
        return nil, err
    }
    for _, order := range GetOrdersRequest.Orders {
        orders = append(orders, service2GraphOrder(order))
    }

  }
  return orders, nil
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

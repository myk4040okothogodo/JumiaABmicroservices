package resolvers

import (
    service_bv1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B"
    "github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/model"
)


func service2GraphOrder(order *service_bv1.Order) *model.Order {
    return &model.Order{
        ID         : order.Id,
        Email      : order.Email,
        Weight     : order.Weight,
        Phonenumber: order.Phonenumber,
        Countrycode: order.Countrycode,
    }
}



func graph2ServiceOrderInput(orderInput *model.OrderInput) *service_bv1.Order {
    return &service_bv1.Order{
        Id:     softDeference(orderInput.ID),
        Email:  softDeference(orderInput.Email),
        Weight: softDeference(orderInput.Weight),
        Phonenumber: softDeference(orderInput.Phonenumber),
        Countrycode: softDeference(orderInput.Countrycode),
    }
}


func softDeference(field string) string {
    if field == nil {
        return ""
    }
    return *field
}

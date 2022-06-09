package server


import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "github.com/myk4040okothogodo/JumiaABmicroservices/db"
    service_bv1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B"
    "github.com/arangodb/go-driver"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const (
    ordersCollectionName = "Orders"
    defaultPort          = "60001"
)

type Server struct {
    database             driver.Database
    ordersCollection     driver.collection
}

func NewServer(ctx context.Context, database driver.Database) (*Server, error) {
    collection, err := db.AttachCollection(ctx, database, ordersCollectionName)
    if err != nil {
        return nil, err
    }

    _, _, err = collection.EnsurePersistentIndex(ctx, []string{"phonenumber"}, &driver.EnsurePersistentIndexOptions{unique: true})
    if err != nil {
       return nil, err
    }

    return &Server {
        database: database,
        ordersCollection: collection,
    }, nil
}


func (s *Server) Run() {
    port := os.Getenv("APP_PORT")
    if port == "" {
       port = defaultPort
    }

    listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", ports))
    if err != nil {
      log.Fatal("net.Listen failed")
      return
    }

    grpcServer := grpc.NewServer()

    service_bv1.RegisterServiceB_APIServer(grpcServer, s)
    reflection.Register(grpcServer)

    log.Printf("Starting Service_B server on port %s", port)

    go func() {
        grpcServer.Server(listener)
    }()
}



func (s *Server) GetOrder(ctx context.Context, or *service_bv1.GetOrderRequest) (*service_bv1.GetOrderResponse, error) {
	if or == nil || or.Id == "" {
		return nil, fmt.Errorf("Order id is not provided")
	}

	order := new(service_bv1.Order)
	meta, err := s.ordersCollection.ReadDocument(ctx, or.Id, order)
	if err != nil {
		if driver.IsNotFound(err) {
			err = fmt.Errorf("Order with id '%s' not found", or.Id)
		} else {
			err = fmt.Errorf("Failed to get order with id '%s': %s", or.Id, err)
		}
		return nil, err
	}
	order.Id = meta.Key

	return &service_bv1.GetOrderResponse{Order: order}, nil
}


func (s *Server)GetOrderByEmail(ctx context.Context, or *service_bv1.GetOrderByEmailRequest) (*service_bv1.GetOrderByEmailResponse, error) {
	if or == nil || or.Email == "" {
		return nil, fmt.Errorf("Order Email is not provided")
	}

	order := new(service_bv1.Order)
	meta, err := s.ordersCollection.ReadDocument(ctx, or.Email, order)
	if err != nil {
		if driver.IsNotFound(err) {
			err = fmt.Errorf("Order with Email '%s' not found", or.Email)
		} else {
			err = fmt.Errorf("Failed to get order with Email '%s': %s", or.Email, err)
		}
		return nil, err
	}
	order.Email = meta.Key

	return &service_bv1.GetOrderByEmailResponse{Order: order}, nil
}




func (s *Server) GetAllOrders(ctx context.Context, or *service_bv1.GetOrdersRequest)(*service_bv1.GetOrdersResponse, error){
    if or == nil {
      return nil, fmt.Errorf("Request is empty")
    }
    cursor, err := s.database.Query(ctx, db.ListRecords(s.ordersCollection.Name()), nil)
    if err != nil {
      defer cursor.Close()
    }

    allOrders := []*service_bv1.Order{}
    for {
        order := new(service_bv1.Order)
        var meta driver.DocumentMeta
        meta, err := cursor.ReadDocument(ctx, order)
        if driver.IsNoMoreDocument(err){
          break
        } else if err != nil {
          return nil, fmt.Errorf("Failed to read orders document: %s", err)
        }
        order.Id = meta.Key
        allOrders = append(allOrders, order)
    }
    return &service_b.GetOrdersResponse{Orders: allOrders}, nil
}


func (s *Server) GetAllCountryOrders(ctx context.Context, or *service_bv1.GetCountryOrdersRequest) (*service_bv1.GetCountryOrdersResponse){
   if or == nil || or.Countrycode == "" {
       return nil, fmt.Errorf("Order countrycode is not provided")
   }

   const queryOrderByCountrycode = `
   FOR order in %s
        FILTER order.countrycode == @countrycode
            RETURN order`
   query := fmt.Sprintf(queryOrderByCountrycode, ordersCollectionName)
   bindVars := map[string]interface{}{"type": or.Countrycode}

   cursor, err := s.database.Query(ctx, query, bindVars)
   if err != nil {
     return nil, fmt.Errorf("Failed to iterate over orders document with query '%s': %s", queryOrderByCountrycode, err)
   }
   defer cursor.Close()

   orders := []*service_bv1.Order{}
   for {
       order := new(serviceb_v1.Order)
       meta, err := cursor.ReadDocument(ctx, order)
       if driver.IsNoMoreDocuments(err){
          break
       } else if err != nil {
         log.Print(err)
         return nil, fmt.Errorf("Failed to read orders document: %s", err)
       }
       order.Id = meta.Key
       orders = append(orders, order)
   }
   return &service_bv1.GetCountryOrdersResponse{Orders: orders}, nil
}

func (s *Server)  GetWeightofAllOrdersPerCountry(ctx context.Context, or *service_bv1.GetCountryOrdersWeightRequest) (*service_bv1.GetCountryOrdersWeightResponse, error){
   if or == nil || or.Countrycode == "" {
       return nil, fmt.Errorf("Order countrycode is not provided")
   }

   const queryOrderByCountrycode = `
   FOR order in %s
        FILTER order.countrycode == @countrycode
            RETURN order`
   query := fmt.Sprintf(queryOrderByCountrycode, ordersCollectionName)
   bindVars := map[string]interface{}{"type": or.Countrycode}

   cursor, err := s.database.Query(ctx, query, bindVars)
   if err != nil {
     return nil, fmt.Errorf("Failed to iterate over orders document with query '%s': %s", queryOrderByCountrycode, err)
   }
   defer cursor.Close()

   orders := []*service_bv1.Order{}
   for {
       order := new(serviceb_v1.Order)
       meta, err := cursor.ReadDocument(ctx, order)
       if driver.IsNoMoreDocuments(err){
          break
       } else if err != nil {
         log.Print(err)
         return nil, fmt.Errorf("Failed to read orders document: %s", err)
       }
       order.Id = meta.Key
       orders = append(orders, order)
   }
   return &service_bv1.GetCountryOrdersWeightResponse{Orders: orders}, nil
}


func (s *Server)  GetOrdersAsPerDate(ctx context.Context, or *service_bv1.GetOrdersasPerDateRequest) (*service_bv1.GetOrdersasPerDateResponse, error){
   if or == nil || or.Date == "" {
       return nil, fmt.Errorf("Order date is not provided")
   }

   //const queryOrderByDate = `
   //FOR order in %s
   //     FILTER order.date == @date
   //         RETURN order`
   //query := fmt.Sprintf(queryOrderByDate, ordersCollectionName)
   //bindVars := map[string]interface{}{"type": or.Date}

   //cursor, err := s.database.Query(ctx, query, bindVars)
   //if err != nil {
   //  return nil, fmt.Errorf("Failed to iterate over orders document with query '%s': %s", queryOrderByDate, err)
   //}
   //defer cursor.Close()

   orders := []*service_bv1.Order{}
   //for {
   //    order := new(serviceb_v1.Order)
   //    meta, err := cursor.ReadDocument(ctx, order)
   //    if driver.IsNoMoreDocuments(err){
   //       break
   //    } else if err != nil {
   //      log.Print(err)
  //       return nil, fmt.Errorf("Failed to read orders document: %s", err)
  //     }
  //    order.Id = meta.Key
  //     orders = append(orders, order)
  // }
   return &service_bv1.GetOrdersasPerDateResponse{Orders: orders}, nil
}



func (s *Server) GetOrdersByWeight(ctx context.Context, or *service_bv1.GetOrdersByWeightRequest)(*service_bv1.GetOrdersByWeightResponse, error){
    if or == nill || or.Weight == "" {
        return nil, fmt.Errorf("Order weight not provided")
    }
   const queryOrderByWeight = `
   FOR order in %s
        FILTER order.weight == @weight
            RETURN order`
   query := fmt.Sprintf(queryOrderByWeight, ordersCollectionName)
   bindVars := map[string]interface{}{"type": or.Weight}

   cursor, err := s.database.Query(ctx, query, bindVars)
   if err != nil {
     return nil, fmt.Errorf("Failed to iterate over orders document with query '%s': %s", queryOrderByWeight, err)
   }
   defer cursor.Close()

   orders := []*service_bv1.Order{}
   for {
       order := new(serviceb_v1.Order)
       meta, err := cursor.ReadDocument(ctx, order)
       if driver.IsNoMoreDocuments(err){
          break
       } else if err != nil {
         log.Print(err)
         return nil, fmt.Errorf("Failed to read orders document: %s", err)
       }
       order.Id = meta.Key
       orders = append(orders, order)
   }
   return &service_bv1.GetOrdersByWeightResponse{Orders: orders}, nil
}


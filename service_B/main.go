package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/myk4040okothogodo/JumiaABmicroservices/service_B/server"
    "github.com/myk4040okothogodo/JumiaABmicroservices/db"
    service_av1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_A"
    service_bv1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B"
    "google.golang.org/grpc"
)


func main(){

  ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
  defer cancelFn()
 
  conn,err := grpc.Dial("localhost:60002", grpc.WithInsecure())
  if err != nil {
    log.Printf("Couldnt connect with  'service_A' service :%s", err)
  }
  defer conn.Close()


  //create client
  sa := service_av1.NewServiceA_APIClient(conn)


  database, err := db.Connect(ctx, db.GetDbConfig())
  if err != nil {
      log.Fatalf("d.OpenDatabase failed with error: %s", err)
  }

  srv, err := server.NewServer(ctx, database, sa)
  if err != nil {
     log.Fatalf("NewServer failed with error: %s", err)
  }

  srv.Run()
  srv.PopulateDatabaseWithOrders(ctx, &service_bv1.PopulateDatabaseWithOrdersRequest{})
  if err != nil {
      log.Printf("Couldnt load the recent Orders :%s", err)
  }


  sigChan := make(chan os.Signal, 1)
  signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
  signal := <-sigChan
  log.Printf("Shutting down Service_B server with signal: %s", signal)

}


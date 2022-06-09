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
)


func main(){

  ctx, cancelFn := context.WithTimeout(context.Baackground(), time.Second*5)
  defer cancelFn()

  database, err := db.Connect(ctx, db.GetDbConfig())
  if err != nil {
      log.Fatalf("d.OpenDatabase failed with error: %s", err)
  }

  srv, err := server.NewServer(ctx, database)
  if err != nil {
     log.Fatalf("NewServer failed with error: %s", err)
  }

  srv.Run()

  sigChan := make(chan os.Signal, 1)
  signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
  signal := <-sigChan
  log.Printf("Shutting down Service_B server with signal: %s", signal)

}

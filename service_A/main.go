package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/myk4040okothogodo/JumiaABmicroservices/service_A/server"
)



func main(){
    ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
    defer cancelFn()

    srv, err := server.NewServer(ctx)
    if err != nil {
        log.Fatalf("NewServer failed with error: %s", err)
    }

    srv.Run()

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    signal := <-sigChan
    log.Printf("Shutting down Service_A server with signal: %s", signal)
}

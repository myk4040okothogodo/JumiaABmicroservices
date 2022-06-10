package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/services"
    "github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/resolvers"
    "github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/gen"
    "github.com/joho/godotenv"
    "github.com/99designs/gqlgen/graphql/handler"
	  "github.com/99designs/gqlgen/graphql/playground"
    "github.com/gorilla/mux"
)

const defaultPort = "8080"
const  defaultServiceBSvc = "localhost:60001"
func parseEnvVars(key string) string {
    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
        log.Printf("Error loading .env file %s", err)
    }

    return os.Getenv(key)
}


func main() {
    //Load environment variables
    port := parseEnvVars("GRAPH_API_PORT")
    if port == "" {
        log.Printf("Failed to load 'GRAPH_API_PORT'(you can set manually) system is reverting to default %s", defaultPort)
        port = defaultPort
    }

    serviceBSvc := os.Getenv("SERVICEB_SERVICE")
    if serviceBSvc == "" {
        log.Printf("Failed to load 'SERVICEB_SERVICE'(you can set manually) system is reverting to default %s", defaultServiceBSvc )
        serviceBSvc = defaultServiceBSvc
    }


    //connect to the services
    svc , err := services.NewServicesKeeper(services.ServicesConfig{
      ServiceBSvc : serviceBSvc,
    })
    if err != nil {
        log.Fatalf("Failed to create grpc api holder: %s", err)
    }


    // Create graphAPI handlers
    router := mux.NewRouter()
    graphAPIHandler := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: resolvers.NewResolver(svc)}))
    router.Handle("/", playground.Handler("GraphQL playground", "/jumiaABmicroservices"))
    router.Handle("/jumiaABmicroservices", graphAPIHandler)


    srv := &http.Server{
        Addr:    fmt.Sprintf(":%s", port),
        WriteTimeout: time.Second * 10,
        ReadTimeout:  time.Second * 10,
        Handler     :  router,
    }


    // start the graph_api server
    log.Printf("Please connect to http://localhost:%s for GrapQL playground", port)
    go func(){
        log.Fatal(srv.ListenAndServe())
    }()

    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)


    //Block untill cancel signal is received
    <-sig
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    log.Print("Shutting down JumiaGraphql GateWay server")

    if err := srv.Shutdown(ctx); err != nil {
        log.Print(err)
    }

    <-ctx.Done()
    os.Exit(0)
}

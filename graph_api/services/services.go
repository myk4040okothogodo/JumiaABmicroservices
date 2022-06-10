package services


import (
    "io"
    "log"
    service_bv1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B"
    "google.golang.org/grpc"
)

type ServicesConfig struct {
   ServiceBSvc  string
}

type services struct {
    io.Closer
    serviceBClientConn   *grpc.ClientConn
    serviceBClient       service_bv1.ServiceB_APIClient

}

type Services interface {
    ServiceB()   service_bv1.ServiceB_APIClient
}


func NewServicesKeeper(conf ServicesConfig)(Services, error) {
    log.Printf("Connecting to ServiceB service at : %s...", conf.ServiceBSvc)
    servicebConnection, err := grpc.Dial(conf.ServiceBSvc, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    sh := &services{
        serviceBClientConn :   servicebConnection,
        serviceBClient     :   service_bv1.NewServiceB_APIClient(servicebConnection),
    }
    return sh, nil
}

func (sh *services)  ServiceB() service_bv1.ServiceB_APIClient {
    return sh.serviceBClient
}



func (sh *services) Close() error {
    err := sh.serviceBClientConn.Close()
    if err != nil {
        log.Printf("An error occurred while closing connection on ServiceB service: %s", err)
    }
    return nil
}

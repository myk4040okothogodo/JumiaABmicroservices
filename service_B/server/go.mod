module github.com/myk4040okothogodo/JumiaABmicroservices/service_B

go 1.18

replace github.com/myk4040okothogodo/JumiaABmicroservices/db => ../../db

replace github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B => ../../gen/go/service_B

require (
	github.com/arangodb/go-driver v1.3.2
	github.com/myk4040okothogodo/JumiaABmicroservices/db v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.47.0
)

require (
	github.com/arangodb/go-velocypack v0.0.0-20200318135517-5af53c29c67e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

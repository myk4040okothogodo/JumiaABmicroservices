module github.com/myk4040okothogodo/JumiaABmicroservices/service_A

go 1.18

replace github.com/myk4040okothogodo/JumiaABmicroservices/service_A/server => ./server

require github.com/myk4040okothogodo/JumiaABmicroservices/service_A/server v0.0.0-20220611070458-2e931f87419f //v0.0.0-00010101000000-000000000000

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_A v0.0.0-20220609205855-b114a567b7aa // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

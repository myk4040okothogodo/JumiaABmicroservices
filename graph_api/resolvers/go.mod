module github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/resolvers

go 1.18

replace github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B => ../../gen/go/service_B

replace github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/model => ../model

replace github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/services => ../services

replace github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/gen => ../gen

require (
	github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_B v0.0.0-20220610071312-0375c1474ac5
	github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/gen v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/model v0.0.0-20220610071312-0375c1474ac5
	github.com/myk4040okothogodo/JumiaABmicroservices/graph_api/services v0.0.0-00010101000000-000000000000
)

require (
	github.com/99designs/gqlgen v0.17.9 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/vektah/gqlparser/v2 v2.4.4 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

// Run command line in root project,

in dummy.proto: option go_package = "goproto/greet";
> protoc --go_out=. --go-grpc_out=. path.proto (proto/dummy.proto)
result: goproto/greet/dummy.pb.go&dummy_grpc.pb.go

in dummy.proto: option go_package = "proto";
> protoc --go_out=. --go-grpc_out=. path.proto (proto/dummy.proto)
result: proto/dummy.pb.go&dummy_grpc.pb.go

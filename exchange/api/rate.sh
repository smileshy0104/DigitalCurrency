protoc rate.proto --go_out=./types --go-grpc_out=./types

goctl rpc protoc rate.proto --go_out=./types --go-grpc_out=./types --zrpc_out=./common --style go_zero

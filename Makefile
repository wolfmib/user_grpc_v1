build:
	protoc -I proto proto/user_proto.proto --go_out=plugins=grpc:user_proto
	GOOS=linux GOARCH=amd64 go build -o user_grpc_v1 && docker build -t user_grpc_v1 .



run: 
	docker run -p 5001:5001 -e MICRO_SERVER_ADDRESS=:5001 -e MICRO_REGISTRY=mdns user_grpc_v1

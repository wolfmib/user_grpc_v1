build:
	...
	GOOS=Linux GOARCH=amd64 go build
	docker build -t user_grpc_v1 .

run:
	docker run -p 5001:5001 user_grpc_v1

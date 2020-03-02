FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD user_grpc_v1 /app/user_grpc_v1

CMD ["./user_grpc_v1"]

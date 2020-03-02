FROM golang:1.14.0-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir app
WORKDIR /app



COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user_grpc_v1



ENV PORT 5001
EXPOSE 5001
CMD ["go", "run"]



FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/user_grpc_v1 .

CMD ["./user_grpc_v1"]

## SOP

---

mkdir user_proto
protoc -I proto proto/user_proto.proto --go_out=plugins=grpc:user_proto

go run main.go 


---

## Easy shell tool
---
source ja_create_golang_env
source ja_grpc_create_user_proto
go run main.go

---

## Test
---
test the communication ok now ✅
---
go run main.go (Running server)
test api_register by pythoon3 (Client)

---
Result:
![Result](img/test_01_communication_run_server_ok.jpg)

--- 

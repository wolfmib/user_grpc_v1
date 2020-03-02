docker images
docker build -t user_grpc_v1 .
docker run -d -p 5001:5001 user_grpc_v1
docker run golang go get -v github.com/wolfmib/user_grpc_v1

- check the last containner that was executed.
---
docker ps -lq
---


## Remove images with grep "key_words"
---
- docker images -f dangling=true | grep "none" | awk '{print $3}' | xargs docker rmi
---

## Remove all images
---
- docker images -a
- docker rmi $(docker images -a -q)
- manually remove repository:tag , if facing deleteing fail via depencdency issue ...
    - docker rmi user_grpc_v1:latest


## Remove all docker ps
---
- show list:
    - docker ps -a
- remove:
    - docker ps -a | awk '{print $1}' | xargs docker rm make run

## Docker imges build and run the container
---
- files
    - Dockerfile
    - Makefile

---
go get -u github.com/golang/protobuf/protoc-gen-go
make build
make run



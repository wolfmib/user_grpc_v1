package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	//ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempey"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempey"`
}

// Global client
var client *mongo.Client

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var my_user User

	/*
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	*/
	log.Println("my user")
	log.Println(my_user)
	log.Println("---------------")

	err := json.NewDecoder(request.Body).Decode(&my_user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("[INFO]: Get the request")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("test_user").Collection("test_user_collection")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("my user")
	log.Println(my_user)
	result, _ := collection.InsertOne(ctx, my_user)
	json.NewEncoder(response).Encode(result)
}

func main() {

	fmt.Println("Start the application ....")

	//Router
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")

	http.ListenAndServe(":12345", router)

}

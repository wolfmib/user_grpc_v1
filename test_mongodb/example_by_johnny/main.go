package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
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

	log.Println(request.Body)

	//custom_user_map := make(map[string]interface{})
	//custom_user_map["firstname"] = "This is interface"
	//custom_user_map["lastname"] = "hahahaha"

	err := json.NewDecoder(request.Body).Decode(&my_user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(my_user)
	// Johnny: this interface, will help you to dynamic access the api data format
	// no need to limit the structure in struct.

	log.Println("[INFO]: Get the request")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("test_user").Collection("test_user_collection")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("my user")
	log.Println(my_user)

	// Check the data is not the same

	var _key string = "firstname"
	var _value string = my_user.FirstName

	// cursor, err := collection.Find(ctx, bson.M{"firstname": "hi"})
	cursor, err := collection.Find(ctx, bson.M{_key: _value})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	defer cursor.Close(ctx)

	var _tem_user User
	for cursor.Next(ctx) {

		log.Println("####### message  Duplicate Data ...##############")
		log.Println("Input firstname", _value)
		cursor.Decode(&_tem_user)
		fmt.Println(_tem_user)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "Duplicate Data ..."} `))

		return
	}

	log.Println("Saveing data ...")
	result, _ := collection.InsertOne(ctx, my_user)
	json.NewEncoder(response).Encode(result)
}

func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {

	// Header
	response.Header().Add("content-type", "application/json")

	// Conn
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("test_user").Collection("test_user_collection")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	// Query All
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": err for query by find()}`))
		return
	}

	defer cursor.Close(ctx)

	var all_users []User
	for cursor.Next(ctx) {
		var _tem_user User

		cursor.Decode(&_tem_user)
		all_users = append(all_users, _tem_user)
	}

	json.NewEncoder(response).Encode(all_users)

}

func main() {

	fmt.Println("Start the application ....")

	//Router
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")

	http.ListenAndServe(":12345", router)

}

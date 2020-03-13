package main

// GateClient  call -> User Services
import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/wolfmib/user_grpc_v1/user_proto"
)

// [GRPC][Johnny]: GRPC Part Constant
const (
	address = "localhost:5001"
)

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name  string             `json:"first_name,omitempty" bson:"first_name,omitempey"`
	Family_Name string             `json:"family_name,omitempty" bson:"family_name,omitempey"`
	Email       string             `json:"email,omitempty" bson:"email,omitempey"`
}

// Johnny:
// This prevent the _ID (000000) dupicate saving into db and caused error
type Create_User struct {
	First_Name  string `json:"first_name,omitempty" bson:"first_name,omitempey"`
	Family_Name string `json:"family_name,omitempty" bson:"family_name,omitempey"`
	Email       string `json:"email,omitempty" bson:"email,omitempey"`
}

// Global client
var client *mongo.Client

//Regiser
func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var my_user pb.RegisterRequest

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

	// Johnny: Call DB for testing, now usring GRPC
	/*
		log.Println("[INFO]: Get the request")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		collection := client.Database("user_db").Collection("user_collection")
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		log.Println("my user")
		log.Println(my_user)

		log.Println("Saveing data ...")
		result, _ := collection.InsertOne(ctx, my_user)
		json.NewEncoder(response).Encode(result)
		logrus.Info(result)
	*/

	//####################################### GRPC #############################
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	r, err := client.RegisterApi(context.Background(), &my_user)

	if err != nil {
		logrus.Error("Could not create user", err)
		logrus.Error(r)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": err for query by find()}`))
		return
	}

	logrus.Info("Created", r)

	//##########################################################################

}

// Query all users
func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {

	// Header
	response.Header().Add("content-type", "application/json")

	// Conn
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("user").Collection("user_collection")
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

// Query one specific user by id
func CetUser_by_id_Endpoint(response http.ResponseWriter, request *http.Request) {

	// Header
	response.Header().Add("content-type", "application/json")

	// Make user structure instance
	var user User

	// load dynamic endpoint request to parameters
	user_dynamic_parameter := mux.Vars(request)

	// Convert to mongo object id (I won't use that )
	id, _ := primitive.ObjectIDFromHex(user_dynamic_parameter["id"])
	log.Println("Get id:    , ", id)

	// Setting ctx
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("test_user").Collection("test_user_collection")

	// Query ici
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	// Drror checking
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": err for query by find()}`))
		return
	}

	// Make json response
	log.Println("==================")
	log.Println("Return the query message user: ", user)
	log.Println("==================")
	json.NewEncoder(response).Encode(user)

}

// Query one specific user by firstname
// /user/name/{firstname}
func CetUser_by_name_Endpoint(response http.ResponseWriter, request *http.Request) {

	// Header
	response.Header().Add("content-type", "application/json")

	// Make user structure instance
	var user User

	// load dynamic endpoint request to parameters
	user_dynamic_parameter := mux.Vars(request)

	// Convert to mongo object id (I won't use that )
	firstname := user_dynamic_parameter["firstname"]
	log.Println("Get firstname:    , ", firstname)

	// Setting ctx
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("test_user").Collection("test_user_collection")

	// Query ici
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"firstname": firstname}).Decode(&user)

	// Drror checking
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": err for query by find()}`))
		return
	}

	// Make json response
	log.Println("==================")
	log.Println("Return the query message user: ", user)
	log.Println("==================")
	json.NewEncoder(response).Encode(user)

}

func main() {

	logrus.Info("[GATE]: Start the application ....")
	logrus.Warn("Use the api_post_create_user.sh to call my /register endpoint !")

	//Router
	router := mux.NewRouter()
	router.HandleFunc("/register", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")                         //Usersssssssss : all user !
	router.HandleFunc("/user/{id}", CetUser_by_id_Endpoint).Methods("GET")               //User, one user with id
	router.HandleFunc("/user/name/{firstname}", CetUser_by_name_Endpoint).Methods("GET") //User, one user with firstanme
	http.ListenAndServe(":12345", router)

}

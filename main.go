package main

//user_grpc_v1
// User Services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/wolfmib/user_grpc_v1/user_proto"
	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func get_mongo_client(INPUT_TIME time.Duration, INPUT_APPLY_URI string) *mongo.Client {

	logrus.Warn("[get_mongo_client]Put me in the 'ja_golang_db package'")
	logrus.Warn("[Jean]: Tu peut faire ja.golang_db.get_client return  mongodb_client_instance ou postgres_client_instance")
	//Setting timeout : 10 seconds :
	// 10 * time.Secound
	local_ctx, _ := context.WithTimeout(context.Background(), INPUT_TIME*time.Second)
	// Get client by url:
	// .ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.Connect(local_ctx, options.Client().ApplyURI(INPUT_APPLY_URI))
	if err != nil {
		err_str := " Can't make connectin to mongodb ... "
		logrus.Error(err_str)
		log.Fatal(err_str)
	}
	return client
}

// mongo.InsertOne Function
func mongo_create_method(INPUT_TIME time.Duration, INPUT_MAP map[string]interface{}, INPUT_COLLECTION *mongo.Collection) (*mongo.InsertOneResult, error) {

	logrus.Warn("[mongo_create_method]Put me in the 'ja_golang_db package'")
	local_ctx, _ := context.WithTimeout(context.Background(), INPUT_TIME*time.Second)

	// [Jean]: bson.M{} is just named type for map[string]interface{}
	//         as you can see in docs: http://godoc.org/labix.org/v2/mgo/bson#M

	/****************************************************************
	custom_bsonM_query := bson.M{}

	for key , value := range INPUT_MAP{
		custom_bsonM_query[key] = value
	}
	****************************************************************/

	custom_bsonM_query, _ := map_to_bsonM(INPUT_MAP)

	// Refernced Code:
	// local_res, err := INPUT_COLLECTION.InsertOne(local_ctx,bson.M{"first_name": "mongo_create_method","family_name":"testing..."})
	local_res, err := INPUT_COLLECTION.InsertOne(local_ctx, custom_bsonM_query)

	id := local_res.InsertedID
	if err != nil {
		logrus.Warn("Can't insert the data")
	} else {
		logrus.Info("Insert one row with ID:  ", id)
	}
	return local_res, err

}

/* [Jason]: the reason I need is
    - for map -> bsonM -> MongoDB instance
	- for map -> xxxxx -> postgres instance
	- so i can have structure:
		- map format ->  insert_any_db_interface -> mongodb_insert_function(map)
		- map format ->  insert_any_db_interface -> postgres_insert_function(map)
*/

func map_to_bsonM(my_map map[string]interface{}) (bson.M, error) {
	logrus.Warn("[map_to_bsonM]: Put me in the 'ja_golang_db package'")

	_tem_bsonM := bson.M{}
	for key, value := range my_map {
		_tem_bsonM[key] = value
	}
	return _tem_bsonM, nil
}

//func request_to_map()

const (
	port = ":5001"
)

//globaly variable for mongodb collection
var user_collection_global *mongo.Collection

// server is used to implement helloworld.
type server struct {
	pb.UnimplementedUserServiceServer
}

// RegisterApi implements UserServicesServer
func (s *server) RegisterApi(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	log.Printf("Received: %v  ", in.GetFirstName())
	log.Printf("Received: %v  ", in.GetFamilyName())
	log.Printf("Received: %v  ", in.GetEmail())

	// key_format in regiter api
	var _key_first_name string = "first_name"
	var _key_family_name string = "family_name"
	var _key_email string = "email"
	var _key_user_id string = "user_id"

	// [Jean]: Insert Format
	input_data_map := make(map[string]interface{})
	input_data_map[_key_first_name] = in.GetFirstName()
	input_data_map[_key_family_name] = in.GetFamilyName()
	input_data_map[_key_email] = in.GetEmail()
	input_data_map[_key_user_id] = 911

	// [Johnny]: Checking the data is dulpicate or not
	local_ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var _first_name_value string = in.GetFirstName()
	var _family_name_value string = in.GetFamilyName()
	var _email_value string = in.GetEmail()

	cur, err := user_collection_global.Find(local_ctx, bson.M{_key_first_name: _first_name_value, _key_family_name: _family_name_value, _key_email: _email_value})
	if err != nil {
		logrus.Error(err)
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	logrus.Info("----------")
	for cur.Next(context.Background()) {
		logrus.Warn("Got Duplicate Register Rquest", cur)
		//Response Back with Error Code 7575 (Repeat Repeat)
		return &pb.RegisterResponse{Uuid: "", Email: "", UserId: 0, ErrorCode: 7575}, errors.New("Duplicated Data Create Request")
	}
	logrus.Info("-------------")
	logrus.Info("No Duplicate Data Found ! Prepare for creating data")
	// [Johnny]: Insert by custom function
	logrus.Info("-------------")
	res, err := mongo_create_method(5, input_data_map, user_collection_global)
	logrus.Info("-------------")
	if err != nil {
		err_str := "mongo_create_method call fail "
		logrus.Error(err_str)
		log.Fatal(err_str)
	} else {
		logrus.Info("Get Inserted data:\n", res)
	}

	email_str := fmt.Sprintf("%v", input_data_map["email"])
	var user_id_int int32 = 911

	return &pb.RegisterResponse{Uuid: "xxxxx", Email: email_str, UserId: user_id_int}, nil
}

func main() {

	logrus.Info("[Jean]: I am backend... Server Start !")
	logrus.Info("[Jean]: C'est User Services !! ")
	logrus.Info("Conn Mongodb now ...")

	// [Jean]: ParExample: 10s  ,  "mongodb://localhost:27017"
	mongo_client := get_mongo_client(10, "mongodb://localhost:27017")

	/*[Jean]: Accees a user_collection pour une utilisation globale..
	  Apres , vous pouvez l'utiliser dans n'importe laquelle de vos fonctions grpc. */
	user_collection_global = mongo_client.Database("user_db").Collection("user_collection")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}

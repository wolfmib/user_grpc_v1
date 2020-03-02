package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Hello World")

	// Mongo Cliient
	// Ctx with timeout 10 s to conn db
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error: Can't make connecton to mongodb")
	}

	// Mongodb collection
	user_collection := client.Database("user_db").Collection("user_collection")

	// Ctx with timeout 5 s to insert data
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := user_collection.InsertOne(ctx, bson.M{"first_name": "golang", "family_name": "test_mongodb_conn"})
	id := res.InsertedID
	fmt.Printf("Insert id = %v      \n", id)

	// Ctx with timout 10 s to query data
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := user_collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		fmt.Println(result)
		fmt.Println("-------------------------------------------------------")
		fmt.Printf(" The format   =             %v \n", result["first_name"])
		fmt.Printf(" The format   =             %v \n", result["family_name"])
		/*

			fmt.Printf(" The format   =             %v ", result["email"])
			fmt.Printf(" The format   =             %v ", result["user_id"])
			fmt.Printf(" The format   =             %v ", result["_id"])
		*/

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}

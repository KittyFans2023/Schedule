package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Make_db(json_data map[string]map[string]map[int]map[string][]string) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://projectlaim2023:jTpSqRamIKn3UTT2@cluster0.lxtqivz.mongodb.net/").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	client.Database("CFU").Collection("schedule").Drop(context.TODO()) //ощищаем нашу коллекцию с расписанием
	collection := client.Database("CFU").Collection("schedule")
	collection.InsertOne(context.TODO(), json_data) //записываем в нее новое

}

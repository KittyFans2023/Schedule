package db

import (
	"context"
	"encoding/json"
	"schedule/GO/schedule/excel_scrapper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Make_db(data []excel_scrapper.Info) {
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
	for _, e := range data {
		collection.InsertOne(context.Background(), e) //записываем в нее новое
	}
}

func Info_about(group ...string) string {
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
	collection := client.Database("CFU").Collection("schedule")
	var documents []map[string]interface{}
	var cursor *mongo.Cursor
	if group[0] != "all" {
		cursor, err = collection.Find(context.Background(), bson.M{"group": group[0]})

	} else {
		cursor, err = collection.Find(context.Background(), bson.M{})
	}
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.Background()) {
		var document map[string]interface{}
		err := cursor.Decode(&document)
		if err != nil {
			panic(err)
		}
		documents = append(documents, document)
	}
	json_data, err := json.Marshal(documents)
	if err != nil {
		panic(err)
	}

	return string(json_data)
}

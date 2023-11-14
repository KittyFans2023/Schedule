package db

import (
	"context"
	"encoding/json"
	"fmt"
	"schedule/GO/schedule/excel_scrapper"
	"sort"
	"strconv"
	"time"

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

func Info_about(group string, year int, month int, day int) string {
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
	var cursor *mongo.SingleResult

	cursor = collection.FindOne(context.Background(), bson.M{"group": group, "date_day": day, "date_month": month, "date_year": year})
	var document map[string]interface{}
	err = cursor.Decode(&document)
	if err != nil {
		a, _ := json.Marshal("null")
		return string(a)
	}

	json_data, err := json.Marshal(document)
	if err != nil {
		a, _ := json.Marshal("null")
		return string(a)
	}

	return string(json_data)
}

func Next_pair(group string) string {

	timing := []int{8 * 60, 9*60 + 50, 11*60 + 30, 13*60 + 20, 15 * 60, 16*60 + 40, 18*60 + 20, 20 * 60}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://projectlaim2023:jTpSqRamIKn3UTT2@cluster0.lxtqivz.mongodb.net/").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		a, _ := json.Marshal("null")
		return string(a)

	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	current_day := time.Now()
	minute := current_day.Minute() + current_day.Hour()*60
	timing = append(timing, minute)
	sort.Ints(timing)
	pair_number := -1
	for index := 0; index < len(timing); index++ {
		if timing[index] == minute {
			pair_number = index + 1
		}
	}
	if pair_number == 9 {
		pair_number = 2
	}
	collection := client.Database("CFU").Collection("schedule")
	var cursor *mongo.SingleResult
	if pair_number != -1 {
		cursor = collection.FindOne(context.Background(), bson.M{"group": group, "date_day": current_day.Day(), "date_month": current_day.Month(), "date_year": current_day.Year()})
	}

	var document map[string]interface{}
	err = cursor.Decode(&document)
	if err != nil {
		a, _ := json.Marshal("null")
		return string(a)
	}
	address := document["lessons"].(map[string]interface{})[strconv.Itoa(pair_number)]
	fmt.Println(address)

	json_data, err := json.Marshal(address)
	if err != nil {
		a, _ := json.Marshal("null")
		return string(a)
	}

	return string(json_data)
}

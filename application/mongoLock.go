package application

import (
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

var orderm map[int][]string = make(map[int][]string)

func orderSeats(min int, max int, reservedBy string) {
	for i := min; i < max; i++ {
		filter := bson.D{{"seats", i}, {"reserved", false}}
		update := bson.D{
			{"$set", bson.D{{"reserved", true}}},
			{"$inc", bson.D{{"times", 1}}},
			{"$push", bson.D{{"reservedBy", reservedBy}}},
		}
		err := mt.UpdateWithSession("test", "seats", filter, update)
		if err != nil {
			continue
		}
		orderm[i] = append(orderm[i], reservedBy)
	}
	wg.Done()
}

var wg sync.WaitGroup

func DeleteOeder() {
	col := mt.Client.Database("test").Collection("seats")
	for i := 0; i < 1000; i++ {
		col.UpdateOne(ctx, bson.M{"seats": i}, bson.M{"$set": bson.M{"reservedBy": make([]string, 0), "reserved": false, "times": 0}})
	}
}

func createSeats() {
	col := mt.Client.Database("test").Collection("seats")
	for i := 0; i < 1000; i++ {
		col.InsertOne(ctx, bson.D{
			{"train", "A"},
			{"seats", i},
			{"reserved", false},
			{"reservedBy", make([]string, 0)},
			{"times", 0},
		})
	}
}

func ConcurrencyGo() {
	mt.Connect("mongodb://root:1234@golang-mongo-1:27017,golang-mongo-2:27017,golang-mongo-3:27017/?replicaSet=rs0")
	defer mt.Disconnect()
	num := 1000
	wg.Add(5)
	go orderSeats(0, num, "Z")
	go orderSeats(0, num, "Y")
	go orderSeats(0, num, "X")
	go orderSeats(0, num, "E")
	go orderSeats(0, num, "R")
	wg.Wait()
	log.Println("Done")
}

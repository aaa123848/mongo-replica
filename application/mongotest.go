package application

import (
	"log"
)

type PersonDoc struct {
	Name string `bson: "name,omitempty"`
	Age  int    `bson: "age,omitempty"`
	Good bool   `bson: "good,omitempty"`
}

func InsertGo() {
	mt.Connect("mongodb://root:1234@golang-mongo-1:27017,golang-mongo-2:27017,golang-mongo-3:27017/?replicaSet=rs0")
	defer mt.Disconnect()
	p := PersonDoc{
		Name: "eric",
		Age:  12,
		Good: true,
	}
	col := mt.Client.Database("test").Collection("person")
	_, err := col.InsertOne(ctx, p)
	if err != nil {
		log.Println(err)
	}
}

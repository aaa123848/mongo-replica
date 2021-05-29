package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type A struct {
	Name string `json: "name bson: "name""`
	Age  int    `json: "age bson: "age"" `
}

func main() {
	// query := `{"$eq":"last value"}`
	// var bsonMap bson.M
	// // Use the JSON package's Unmarshal() method
	// err := json.Unmarshal([]byte(query), &bsonMap)
	// if err != nil {
	// 	log.Fatal("json. Unmarshal() ERROR:", err)
	// } else {
	// 	fmt.Println("bsonMap:", bsonMap)
	// 	fmt.Println("bsonMap TYPE:", reflect.TypeOf(bsonMap))
	// 	fmt.Println("BSON:", reflect.TypeOf(bson.M{"int field": bson.M{"$gt": 42}}))
	// }
	// res, err := json.Marshal(bsonMap)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
	// fmt.Println(string(res))
	a := bson.M{"a": 1, "b": 2}
	e, res, err := bson.MarshalValue(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println(string(res))
	fmt.Println(e)
}

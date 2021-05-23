package main

import (
	"context"
	"time"
	//"log"
	//"fmt"
	"math/rand"

	//"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    //"go.mongodb.org/mongo-driver/mongo/readpref"
)


const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func String(length int) string {
	return StringWithCharset(length, charset)
}
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
  
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
  
type MongoTool struct {
	Context context.Context
	Client *mongo.Client
}

func (m *MongoTool) Connect(url string){
	client, err := mongo.Connect(m.Context, options.Client().ApplyURI(url))
	if err != nil{
		panic(err)
	}
	m.Client = client
}

func (m *MongoTool) Disconnect(){
	if err := m.Client.Disconnect(m.Context); err != nil {
		panic(err)
	}
}


var mt MongoTool = MongoTool{}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mt.Context = ctx
	defer cancel()
	mt.Connect("mongodb://root:1234@golang-mongo-1:27017,golang-mongo-2:27017,golang-mongo-3:27017/?replicaSet=rs0")
	defer mt.Disconnect()
	db := mt.Client.Database("test")
	db.CreateCollection(ctx, "lala")	
}
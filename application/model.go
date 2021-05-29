package application

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelInter interface {
	ReadOne(context.Context) interface{}
	ReadAll(context.Context) interface{}
}

type PersonDoc struct {
	Name string `bson: "name,omitempty"`
	Age  int    `bson: "age,omitempty"`
	Good bool   `bson: "good,omitempty"`
}

func (p PersonDoc) getColl() *mongo.Collection {
	return mt.Client.Database("test").Collection("person")
}

func (p PersonDoc) ReadOne(ctx context.Context) interface{} {
	col := p.getColl()
	col.FindOne(ctx, bson.M{}).Decode(&p)
	return p
}

func (p PersonDoc) ReadAll(ctx context.Context) interface{} {
	col := p.getColl()
	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	res := make([]PersonDoc, -0, 50)
	for cursor.Next(ctx) {
		tmp := PersonDoc{}
		cursor.Decode(&tmp)
		res = append(res, tmp)
	}
	return res
}

package application

import (
	"context"
	"errors"
	"log"
	"mongotest/mongotool"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonDoc struct {
	Name string `bson: "name,omitempty"`
	Age  int    `bson: "age,omitempty"`
	Good bool   `bson: "good,omitempty"`
}

func (p PersonDoc) getColl() *mongo.Collection {
	return mt.Client.Database("test").Collection("person")
}

func (p PersonDoc) Varify() error {
	if p.Age > 15 {
		return errors.New("age too big")
	}
	return nil
}

func (p PersonDoc) InsertOne(ctx context.Context, mt mongotool.MongoTool) error {
	col := p.getColl()
	_, err := col.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	return nil
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

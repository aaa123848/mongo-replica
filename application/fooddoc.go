package application

import (
	"context"
	"mongotest/mongotool"

	"go.mongodb.org/mongo-driver/mongo"
)

type FoodDoc struct {
	Color  string `json: "color"`
	Flavor string `json: "flavor"`
}

func (f FoodDoc) getColl(mt mongotool.MongoTool) *mongo.Collection {
	return mt.Client.Database("test").Collection("food")
}

func (f FoodDoc) Varify() error {
	return nil
}

func (f FoodDoc) InsertOne(ctx context.Context, mt mongotool.MongoTool) error {
	col := f.getColl(mt)
	_, err := col.InsertOne(ctx, f)
	if err != nil {
		return err
	}
	return nil
}

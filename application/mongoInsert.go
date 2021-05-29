package application

import (
	"context"
	"fmt"
	"log"
	"mongotest/mongotool"

	"go.mongodb.org/mongo-driver/mongo"
)

type InsertInter interface {
	Varify() error
	InsertOne(ctx context.Context, mt mongotool.MongoTool) error
}

func InsertOne(ctx context.Context, mt mongotool.MongoTool, ins InsertInter) error {
	err := ins.Varify()
	if err != nil {
		return err
	}
	err = ins.InsertOne(ctx, mt)
	if err != nil {
		return nil
	}
	return nil
}

func InsertMany(ctx context.Context, mt mongotool.MongoTool, inss []InsertInter) error {
	session, err := mt.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}
		for _, s := range inss {
			err := s.Varify()
			if err != nil {
				_ = session.AbortTransaction(ctx)
				return err
			}
			err = s.InsertOne(sc, mt)
			if err != nil {
				_ = session.AbortTransaction(ctx)
				return err
			}
			log.Printf("[Info] Success %v Insert\n", s)
		}
		session.CommitTransaction(ctx)
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func InsertGo() {
	mt.Connect("mongodb://root:1234@golang-mongo-1:27017,golang-mongo-2:27017,golang-mongo-3:27017/?replicaSet=rs0")
	defer mt.Disconnect()
	ps := []InsertInter{
		PersonDoc{
			Name: "eee",
			Age:  12,
			Good: true,
		},
		FoodDoc{
			Color:  "red",
			Flavor: "sweet",
		},
		PersonDoc{
			Name: "fff",
			Age:  19,
			Good: true,
		},
	}

	err := InsertMany(ctx, mt, ps)
	if err != nil {
		fmt.Println(err)
	}
}

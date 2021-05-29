package application

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func GetOne(ctx context.Context, m ModelInter) (string, error) {
	q := m.ReadOne(ctx)
	res, err := json.Marshal(q)
	if err != nil {
		log.Println(err)
	}
	return string(res), nil
}

func GetAll(ctx context.Context, m ModelInter) (string, error) {
	q := m.ReadAll(ctx)
	res, err := json.Marshal(q)
	if err != nil {
		log.Println(err)
	}
	return string(res), nil
}

func QueryGo() {
	ctx := context.Background()
	mt.Connect("mongodb://root:1234@golang-mongo-1:27017,golang-mongo-2:27017,golang-mongo-3:27017/?replicaSet=rs0")
	p := PersonDoc{}
	res, err := GetAll(ctx, p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))
}

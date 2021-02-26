package mongo

import (
	"context"

	"github.com/xm-chentl/go-dbfty/grammar"
	"github.com/xm-chentl/go-dbfty/metadata"
	"go.mongodb.org/mongo-driver/mongo"
)

type add struct {
	db      *mongo.Database
	data    interface{}
	grammar grammar.IInsert
}

func (a *add) Exec() error {
	table, err := metadata.Get(a.data)
	if err != nil {
		return err
	}

	// todo.解析结构体为 map
	dataMap := table.GetValueByMap(a.data)
	_, err := a.db.Collection(table.Name()).InsertOne(context.Background(), dataMap)
	if err != nil {
		return err
	}
}

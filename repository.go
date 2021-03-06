package mongo

import (
	"context"
	"fmt"

	"github.com/xm-chentl/go-dbfty"
	"github.com/xm-chentl/go-dbfty/grammar"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db        *mongo.Client
	grammar   grammar.IGrammar
	operation Operation
}

func (r repository) Add(entity interface{}) dbfty.IAdd {
	return nil
}

func (r repository) Ping() error {
	var err error
	if err = r.db.Connect(context.Background()); err != nil {
		return fmt.Errorf("mongo connection error (%s)", err.Error())
	}
	if err = r.db.Database(r.operation.DbName); err != nil {
		return fmt.Errorf("database conenction error (%s)", err.Error())
	}

	return nil
}

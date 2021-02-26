package mongo

import (
	"context"

	"github.com/xm-chentl/go-dbfty/grammar"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db      *mongo.Client
	grammar grammar.IGrammar
}

func (r repository) Ping() error {
	return r.db.Connect(context.Background())
}

func ()
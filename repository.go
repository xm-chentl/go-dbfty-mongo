package mongo

import "github.com/xm-chentl/go-dbfty/grammar"

type repository struct{
	db *mongo.Client
	grammar grammar.IGrammar
}

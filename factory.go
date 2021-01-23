package mongo

import (
	"context"
	"github.com/xm-chentl/go-dbfty"
	"go.mongodb.org/mongo-driver/mongo"
)

type factory struct{
	connStr string
	repository *repository
}

func (f factory) Db() dbfty.IRepository{

	return nil
}

func (f factory) Uow() dbfty.IUnitOfWork{
	return nil
}

func (f factory)  getRepository() *repository{
	if f.repository == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI(f.connStr))
		if err != nil {
			panic(err)
		}
		f.repository = &repository{
			db:client,
			grammar:
		}
	}

	return f.repository
}

func New(connStr string) dbfty.IFactory {
	return &factory{
		connStr: connStr,
	}
}
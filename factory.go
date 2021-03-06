package mongo

import (
	"github.com/xm-chentl/go-dbfty"
	gomongo "github.com/xm-chentl/go-dbfty/grammar/nosql/mongo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Operation struct {
	Host     string
	Port     uint16
	DbName   string
	User     string
	Password string
}

func (o *Operation) ConnStr() string {
	return ""
}

type factory struct {
	connStr    string
	operation  Operation
	repository *repository
}

func (f factory) Db() dbfty.IRepository {
	return f.getRepository()
}

func (f factory) Uow() dbfty.IUnitOfWork {
	return nil
}

func (f factory) IsHealth() error {
	if err := f.getRepository().Ping(); err != nil {
		return err
	}

	return nil
}

func (f factory) getRepository() *repository {
	if f.repository == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI(f.operation.ConnStr()))
		if err != nil {
			panic(err)
		}

		f.repository = &repository{
			db:        client,
			grammar:   gomongo.New(),
			operation: f.operation,
		}
	}

	return f.repository
}

func New(operation Operation) dbfty.IFactory {
	return &factory{
		operation: operation,
	}
}

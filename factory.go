package mongo

import (
	"github.com/xm-chentl/go-dbfty"
	gomongo "github.com/xm-chentl/go-dbfty/grammar/nosql/mongo"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type factory struct {
	connStr    string
	repository *repository
}

func (f factory) Db() dbfty.IRepository {
	return f.getRepository()
}

func (f factory) Uow() dbfty.IUnitOfWork {
	return nil
}

func (f factory) getRepository() *repository {
	if f.repository == nil {
		client, err := mongo.NewClient(options.Client().ApplyURI(f.connStr))
		if err != nil {
			panic(err)
		}
		f.repository = &repository{
			db:      client,
			grammar: gomongo.New(),
		}
	}

	return f.repository
}

func (f factory) IsHealth() (bool, error) {
	return false, nil
}

func New(connStr string) dbfty.IFactory {
	return &factory{
		connStr: connStr,
	}
}

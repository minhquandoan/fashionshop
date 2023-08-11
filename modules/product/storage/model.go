package productstorage

import "go.mongodb.org/mongo-driver/mongo"

type dbStore struct {
	collection *mongo.Collection
}

func NewDbStore(collection *mongo.Collection) *dbStore {
	return &dbStore{collection: collection}
}
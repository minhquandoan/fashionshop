package shopstorage

import "go.mongodb.org/mongo-driver/mongo"

type ShopStorage struct {
	collection *mongo.Collection
}

func NewShopStorage(collection *mongo.Collection) *ShopStorage {
	return &ShopStorage{collection: collection}
}
package userstorage

import "go.mongodb.org/mongo-driver/mongo"

type userStore struct {
	collection *mongo.Collection
}

func NewUserStorage(coll *mongo.Collection) *userStore {
	return &userStore{collection: coll}
}
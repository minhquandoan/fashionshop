package uploadstorage

import "go.mongodb.org/mongo-driver/mongo"

type UploadStorage struct {
	collection *mongo.Collection
}

func NewUploadStorage(collection *mongo.Collection) *UploadStorage {
	return &UploadStorage{collection: collection}
}
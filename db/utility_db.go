package db

import "go.mongodb.org/mongo-driver/mongo"

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(GetDbName()).Collection(collectionName)
}
package shopstorage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (store *ShopStorage) Update(ctx context.Context, id *string, data *bson.D) (interface{}, error) {
	coll := store.collection

	result, err := coll.UpdateByID(ctx, *id, *data)
	if err != nil {
		return nil, err
	}

	return result, nil 
}
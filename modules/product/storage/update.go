package productstorage

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *dbStore) UpdateOneProduct(ctx context.Context, id *primitive.ObjectID, updateData *bson.M) (interface{}, error) {
	coll := store.collection

	result, err := coll.UpdateOne(ctx, bson.M{"_id": *id}, bson.M{"$set": *updateData})
	if err != nil {
		return nil, common.ErrCannotUpdateEntity("Product", err)
	}

	return result, nil
}
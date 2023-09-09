package shopstorage

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *ShopStorage) Update(ctx context.Context, id *string, data *bson.D) (interface{}, error) {
	coll := store.collection

	result, err := coll.UpdateByID(ctx, *id, *data)
	if err != nil {
		return nil, err
	}

	return result, nil 
}

func (store *ShopStorage) UpdateLikedCount(ctx context.Context, id *primitive.ObjectID, value int16) (interface{}, error) {
	coll := store.collection

	addLike := bson.D{{"$inc", bson.D{{"liked_count", value}}}}
	
	result, err := coll.UpdateOne(ctx, bson.D{{"_id", *id}}, addLike)
	if err != nil {
		return nil, common.ErrCannotUpdateEntity("shop", err)
	}

	return common.SimpleSuccessResponse(result), nil
}
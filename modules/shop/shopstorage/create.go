package shopstorage

import (
	"context"
	"time"

	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *ShopStorage) Create(context context.Context, data *shopmodel.ShopModel) (interface{}, error) {
	coll := store.collection

	timezone := primitive.Timestamp{T: uint32(time.Now().Unix())}
	data.CreatedAt = &timezone
	data.UpdatedAt = &timezone
	data.Status = 1

	id, err := coll.InsertOne(context, data)
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}
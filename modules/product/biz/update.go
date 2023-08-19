package productbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateProductStorage interface {
	UpdateOneProduct(ctx context.Context, id *primitive.ObjectID, updateData *bson.M) (interface{}, error)
}

type updateProductBiz struct {
	store UpdateProductStorage
}

func NewUpdateProductBiz(store UpdateProductStorage) *updateProductBiz {
	return &updateProductBiz{store: store}
}

func(biz *updateProductBiz) UpdateOneProduct(ctx context.Context, id *string, updateData *bson.M) (interface{}, error) {
	// Convert ID from string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	result, err := biz.store.UpdateOneProduct(ctx, &objectID, updateData)
	return result, err
}
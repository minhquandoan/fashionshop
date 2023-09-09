package shopbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateShopStorage interface {
	Update(ctx context.Context, id *string, data *bson.D) (interface{}, error)
	UpdateLikedCount(ctx context.Context, id *primitive.ObjectID, value int16) (interface{}, error)
}

type updateShopBiz struct {
	shopStore UpdateShopStorage
}

func NewUpdateShopBiz(store UpdateShopStorage) *updateShopBiz {
	return &updateShopBiz{shopStore: store}
}

func(biz *updateShopBiz) IncreaseLikedCount(ctx context.Context, id *string) (interface{}, error) {
	// Convert ID from string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	
	return biz.shopStore.UpdateLikedCount(ctx, &objectID, 0)
}
package shopbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
)

type CreateShopStorage interface {
	Create(ctx context.Context, data *shopmodel.ShopModel) (interface{}, error)
}

type CreateShopBiz struct {
	store CreateShopStorage
}

func NewCreateShopBiz(store CreateShopStorage) *CreateShopBiz {
	return &CreateShopBiz{store: store}
}

func (biz *CreateShopBiz) AddShop(ctx context.Context, data *shopmodel.ShopModel) (interface{}, error) {
	store := biz.store

	result, err := store.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
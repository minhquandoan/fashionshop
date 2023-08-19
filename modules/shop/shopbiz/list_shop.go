package shopbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
)

type ListShopStorage interface {
	ListById(ctx context.Context, id *string) (*shopmodel.ShopModel, error)
}

type ListShopBiz struct {
	store ListShopStorage
}

func NewListShopBiz(store ListShopStorage) *ListShopBiz {
	return &ListShopBiz{store: store}
}

func (biz *ListShopBiz) ListShopById(ctx context.Context, id *string) (*shopmodel.ShopModel, error) {
	store := biz.store

	result, err := store.ListById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
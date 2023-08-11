package productbiz

import (
	"context"

	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
)

type ListProductStore interface {
	ListProducts(ctx context.Context, productList *[]productmodel.Product) error
}

type listProductBiz struct {
	store ListProductStore
}

func NewListProductBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store:store}
}

func (biz *listProductBiz) ListProducts(ctx context.Context, productList *[]productmodel.Product) error {
	if err := biz.store.ListProducts(ctx, productList); err != nil {
		return err
	}

	return nil
}
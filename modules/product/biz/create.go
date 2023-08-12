package productbiz

import (
	"context"

	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
)

type CreateProductStore interface {
	CreateOneProduct(ctx context.Context, data *productmodel.Product) (interface{}, error)
}

type createProductBiz struct {
	store CreateProductStore
}

func NewCreateProductBiz(store CreateProductStore) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProduct(ctx context.Context, data *productmodel.Product) (interface{}, error) {
	return biz.store.CreateOneProduct(ctx, data)
}
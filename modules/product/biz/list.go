package productbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	"go.mongodb.org/mongo-driver/bson"
)

type ListProductStore interface {
	ListProducts(ctx context.Context, productList *[]productmodel.Product) error
	ListProductsByConditions(ctx context.Context, filter *bson.M, paging *common.Paging, productList *[]productmodel.Product)  error
}

type listProductBiz struct {
	store ListProductStore
}

func NewListProductBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store:store}
}

func (biz *listProductBiz) ListProducts(ctx context.Context, productList *[]productmodel.Product) error {
	err := biz.store.ListProducts(ctx, productList)
	return err
}

func (biz *listProductBiz) ListProductByFilters(ctx context.Context, filter *bson.M, paging *common.Paging, productList *[]productmodel.Product)  error {
	return biz.store.ListProductsByConditions(ctx, filter, paging, productList)
}
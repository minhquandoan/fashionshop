package productstorage

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (store *dbStore) ListProducts(ctx context.Context, productList *[]productmodel.Product) error {
	coll := store.collection

	productCur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return common.ErrDB(err)
	}

	defer productCur.Close(ctx)

	for productCur.Next(ctx) {
		var product productmodel.Product
		if err := productCur.Decode(&product); err != nil {
			return common.ErrInternal(err)
		}

		(*productList) = append((*productList), product)
	}

	return nil
}
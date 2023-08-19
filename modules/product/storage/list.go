package productstorage

import (
	"context"
	"github.com/minhquandoan/fashionshop/common"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *dbStore) ListProducts(ctx context.Context, productList *[]productmodel.Product) error {
	coll := store.collection

	// Find all documents
	productCur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return common.ErrDB(err)
	}

	defer productCur.Close(ctx)

	// Fetch Data from cursor
	for productCur.Next(ctx) {
		var product productmodel.Product
		if err := productCur.Decode(&product); err != nil {
			return common.ErrInternal(err)
		}

		(*productList) = append((*productList), product)
	}

	return nil
}

func (store *dbStore) ListProductsByConditions(ctx context.Context, filter *bson.M, paging *common.Paging, productList *[]productmodel.Product)  error {
	coll := store.collection

	// Filter Stages
	matchStage := bson.D{{"$match", filter}}

	// Pagination Stages
	limitStage := bson.D{{"$limit", paging.Limit}}
	offSetStage := bson.D{{"$skip", (paging.Page - 1) * paging.Limit}}

	//Sorting Stages
	sortStage := bson.D{{"$sort", bson.D{{"created_at", -1}}}}

	// Execute pipelines
	productCur, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, limitStage, offSetStage, sortStage})
	if err != nil {
		return common.ErrDB(err)
	}
	defer productCur.Close(ctx)

	// fetch data from cursor
	for productCur.Next(ctx) {
		var product productmodel.Product

		if err := productCur.Decode(&product); err != nil {
			return common.ErrInternal(err)
		}
		(*productList) = append((*productList), product)
	}

	return nil
}
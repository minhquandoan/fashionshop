package productstorage

import (
	"context"
	"time"

	"github.com/minhquandoan/fashionshop/common"
	productmodel "github.com/minhquandoan/fashionshop/modules/product/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *dbStore) CreateOneProduct(ctx context.Context, data *productmodel.Product) (interface{}, error) {
	coll := store.collection
	
	// Initialize time and activate object
	timezone := primitive.Timestamp{T: uint32(time.Now().Unix())}
	data.CreatedAt = &timezone
	data.UpdatedAt = &timezone
	data.Status = 1

	result, err := coll.InsertOne(ctx, data)
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result.InsertedID, nil
}
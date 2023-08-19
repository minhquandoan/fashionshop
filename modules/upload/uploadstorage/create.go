package uploadstorage

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
)

func (store *UploadStorage) CreateImage(ctx context.Context, data *common.Image) (interface{}, error) {
	coll := store.collection

	result, err := coll.InsertOne(ctx, data);
	if err != nil {
		// log.Fatal(err)
		return nil, common.ErrCannotCreateEntity("", err)
	}

	return result.InsertedID, nil
}
package uploadstorage

import (
	"context"
	"log"

	"github.com/minhquandoan/fashionshop/common"
)

func (store *UploadStorage) ListImages(
		ctx context.Context, 
		ids []interface{}, 
		moreKeys ...string,
) ([]common.Image, error) {

	coll := store.collection

	curs, err := coll.Find(ctx, ids)
	if err != nil {
		log.Fatal("there is no image with id")
		return nil, common.ErrEntityNotFound("", err)
	}

	var results []common.Image
	if err = curs.All(ctx, &results); err != nil {
		// log.Fatal(err)
		return nil, common.ErrInternal(err)
	}

	return results, nil
}
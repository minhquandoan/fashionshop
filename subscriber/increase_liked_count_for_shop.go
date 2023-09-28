package subscriber

import (
	"context"

	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/shop/shopstorage"
	"github.com/minhquandoan/fashionshop/pubsub"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateLikedCountShop interface {
	GetId() *primitive.ObjectID
	GetValue() int16
}

func IncreaseLikedCountForShop(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Update liked count when user like/unlike shop",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			coll := db.GetCollection(appCtx.GetDbClient(), shopmodel.GetCollectionName())
			data := message.Data().(UpdateLikedCountShop)

			shopStore := shopstorage.NewShopStorage(coll)
			_, err := shopStore.UpdateLikedCount(ctx, data.GetId(), data.GetValue())
			return err
		},
	}
}
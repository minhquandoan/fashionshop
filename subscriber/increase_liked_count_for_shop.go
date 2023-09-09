package subscriber

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/shop/shopstorage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateLikedCountShop interface {
	GetId() primitive.ObjectID
	GetValue() int16
}

func IncreaseLikedCountForShop(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	shopStore := shopstorage.NewShopStorage(db.GetCollection(appCtx.GetDbClient(), shopmodel.GetCollectionName()))
	
	go func() {
		defer common.AppRecover()

		for {
			msg := <- c
			data := msg.Data().(UpdateLikedCountShop)
			id := data.GetId()
			shopStore.UpdateLikedCount(ctx, &id, data.GetValue())
		}
	}()
}
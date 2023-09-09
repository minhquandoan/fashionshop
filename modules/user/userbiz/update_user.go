package userbiz

import (
	"context"
	"log"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"github.com/minhquandoan/fashionshop/pubsub"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserStorage interface {
	UpdateLikedShop(ctx context.Context, ops string, data *usermodel.UserLikeShop) (interface{}, error)
	CheckUserLikeShop(ctx context.Context, filter *usermodel.UserLikeShop) (bool, error)
}

type ShopStorage interface {
	ListById(ctx context.Context, id *string) (*shopmodel.ShopModel, error)
	UpdateLikedCount(ctx context.Context, id *primitive.ObjectID, value int16) (interface{}, error)
}

type UpdateUserBiz struct {
	userStore UpdateUserStorage
	shopStore ShopStorage
	pb pubsub.PubSub
}

func NewUpdateUserBiz(userStore UpdateUserStorage, shopStore ShopStorage, pb pubsub.PubSub) *UpdateUserBiz {
	return &UpdateUserBiz{userStore: userStore, shopStore: shopStore, pb:pb}
}

func(biz *UpdateUserBiz) UserLikeShop(ctx context.Context, data *usermodel.UserLikeShop) (interface{}, error) {
	userStore  := biz.userStore
	shopStore := biz.shopStore

	// Check shop's existence
	shopId := data.LikedShop.Hex()
	_, err := shopStore.ListById(ctx, &shopId)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	// Check user like shop or not
	liked, err := userStore.CheckUserLikeShop(ctx, data)
	if err != nil {
		return nil, err
	}
	
	// If shop was liked by user, remove shopID in user (unlike) - otherwise, add shopID to user
	var ops string
	var value int16
	if liked {
		ops = "$pull"
		value = -1
	}else {
		ops = "$push"
		value = 1
	}
	log.Println("Update Like shop: ", ops)

	// Update liked shop ID in user
	result, err := userStore.UpdateLikedShop(ctx, ops, data)
	if err != nil {
		return nil, err
	}

	// side effect
	_ = biz.pb.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(shopmodel.NewShopUpdateLikedCount(data.LikedShop, value)))

	return result, nil
}
package userbiz

import (
	"context"
	"log"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
)

type UpdateUserStorage interface {
	UpdateLikedShop(ctx context.Context, ops string, data *usermodel.UserLikeShop) (interface{}, error)
	CheckUserLikeShop(ctx context.Context, filter *usermodel.UserLikeShop) (bool, error)
}

type FindShopStorage interface {
	ListById(ctx context.Context, id *string) (*shopmodel.ShopModel, error)
}

type UpdateUserBiz struct {
	userStore UpdateUserStorage
	shopStore FindShopStorage
}

func NewUpdateUserBiz(userStore UpdateUserStorage, shopStore FindShopStorage) *UpdateUserBiz {
	return &UpdateUserBiz{userStore: userStore, shopStore: shopStore}
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
	if liked {
		ops = "$pull"
	}else {
		ops = "$push"
	}
	log.Println("Update Like shop: ", ops)

	// Update liked shop ID in user
	result, err := userStore.UpdateLikedShop(ctx, ops, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
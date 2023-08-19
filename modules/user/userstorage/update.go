package userstorage

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
)

func (store *userStore) UpdateLikedShop(ctx context.Context, ops string, data *usermodel.UserLikeShop) (interface{}, error) {
	coll := store.collection

	result, err := coll.UpdateByID(ctx, data.Id, bson.M{ops: bson.M{"liked_shop": data.LikedShop}})
	if err != nil {
		return nil, common.ErrCannotUpdateEntity("Shop", err)
	}

	return result, nil
}
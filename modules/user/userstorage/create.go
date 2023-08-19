package userstorage

import (
	"context"
	"time"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (store *userStore) CreateUser(ctx context.Context, data *usermodel.User) (interface{}, error) {
	coll := store.collection

	timezone := primitive.Timestamp{T: uint32(time.Now().Unix())}
	data.CreatedAt = &timezone
	data.UpdatedAt = &timezone
	data.Status = 1

	result, err := coll.InsertOne(ctx, data)
	if err != nil {
		return nil, common.ErrCannotCreateEntity("", err)
	}

	return result.InsertedID, nil
}
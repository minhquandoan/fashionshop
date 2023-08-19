package userbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, data *usermodel.User) (interface{}, error)
	FindUser(context context.Context, filters *bson.D) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type createUserBiz struct {
	store CreateUserStore
	hasher Hasher
}

func NewRegisterUserBiz(store CreateUserStore, hasher Hasher) *createUserBiz {
	return &createUserBiz{store: store, hasher: hasher}
}

func (biz *createUserBiz) RegisterUser(ctx context.Context, data *usermodel.User) (interface{}, error) {
	store := biz.store
	hasher := biz.hasher

	filters := bson.D{{"email", *data.Email}}

	user, err := store.FindUser(ctx, &filters)
	if user != nil {
		return nil, common.ErrEntityExisted(*data.Email, err)
	}

	salt := common.GenSalt(50)
	*data.Password = hasher.Hash(*data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	result, err := biz.store.CreateUser(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}


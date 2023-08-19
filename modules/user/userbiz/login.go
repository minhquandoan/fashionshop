package userbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/component/tokenprovider"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
)

type UserLoginStorage interface {
	FindUser(context context.Context, filters *bson.D) (*usermodel.User, error)
}

type userLoginBiz struct {
	AppCtx        component.AppContext
	Store         UserLoginStorage
	TokenProvider tokenprovider.Provider
	Hasher        Hasher
	Expiry        int
}

func NewUserLoginBiz(appCtx component.AppContext, store UserLoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *userLoginBiz {
		return &userLoginBiz{
			AppCtx: appCtx,
			Store: store,
			TokenProvider: tokenProvider,
			Hasher: hasher,
			Expiry: expiry,
		}
}

func NewSubUserLoginBiz(store UserLoginStorage) *userLoginBiz {
	return &userLoginBiz{
		Store: store,
		AppCtx: nil,
		TokenProvider: nil,
		Hasher: nil,
		Expiry: 0,
	}
}

func (biz *userLoginBiz) Login(context context.Context, credentials *usermodel.UserAccount) (*tokenprovider.Token, error) {
	filter := bson.D{{"email", *credentials.Email}}
	//Get User By Email
	user, err := biz.Store.FindUser(context, &filter)
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	//Get Salt from user then hash with Password from request
	hashedPass := biz.Hasher.Hash(*credentials.Password + user.Salt)
	if hashedPass != *user.Password {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	// create payload with user and role
	payload := tokenprovider.TokenPayload{
		UserId: user.Id.String(),
		Role: user.Role,
	}

	accessToken, err := biz.TokenProvider.Generate(payload, biz.Expiry)
	if err != nil {
		return nil, tokenprovider.ErrEncodingToken
	}

	return accessToken, nil
}

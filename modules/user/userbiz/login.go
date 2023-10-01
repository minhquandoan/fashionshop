package userbiz

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
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

func (biz *userLoginBiz) Login(ctx context.Context, credentials *usermodel.UserAccount) (*tokenprovider.Token, error) {
	filter := bson.D{{"email", *credentials.Email}}

	// Create new tracer
	tracer, _ := common.NewTracer(ctx, "jaeger-user.biz.login")

	//Get User By Email
	_, span1 := tracer.Start(ctx, "user.biz.login_find")
	user, err := biz.Store.FindUser(ctx, &filter)
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	span1.End()

	//Get Salt from user then hash with Password from request
	_, span2 := tracer.Start(ctx, "user.biz.login_hashing")
	hashedPass := biz.Hasher.Hash(*credentials.Password + user.Salt)
	if hashedPass != *user.Password {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	span2.End()

	_, span3 := tracer.Start(ctx, "user.biz.login_creating-payload")
	// create payload with user and role
	payload := tokenprovider.TokenPayload{
		UserId: user.Id.String(),
		Role: user.Role,
	}

	accessToken, err := biz.TokenProvider.Generate(payload, biz.Expiry)
	if err != nil {
		return nil, tokenprovider.ErrEncodingToken
	}
	span3.End()

	return accessToken, nil
}

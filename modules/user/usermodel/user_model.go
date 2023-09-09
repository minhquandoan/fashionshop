package usermodel

import (
	"errors"

	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionName = "user"
)

type User struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	common.SqlModel `bson:"inline" json:"inline"`
	FirstName       *string       `bson:"first_name,omitempty" json:"first_name,omitempty" form:"first_name"`
	LastName        *string       `bson:"last_name,omitempty" json:"last_name,omitempty" form:"last_name"`
	Email           *string       `bson:"email,omitempty" json:"email,omitempty" form:"email"`
	Password        *string       `bson:"password,omitempty" json:"password,omitempty" form:"password"`
	Salt            string        `bson:"salt,omitempty" json:"salt,omitempty"`
	Phone           *string       `bson:"phone,omitempty" json:"phone,omitempty" form:"phone"`
	Role            string        `bson:"role,omitempty" json:"role,omitempty" form:"role"`
	Avatar          *common.Image `bson:"avatar,omitempty" json:"avatar,omitempty"`
}

func GetCollectionName() string {
	return CollectionName
}

type UserAccount struct {
	Email    *string `bson:"email,omitempty" json:"email,omitempty" form:"email"`
	Password *string `bson:"password,omitempty" json:"password,omitempty" form:"password"`
}

type UserLikeShopModel struct {
	Id        primitive.ObjectID   `bson:"_id,omitempty"`
	LikedShop []primitive.ObjectID `bson:"liked_shop,omitempty" json:"liked_shop,omitempty"`
}

type UserLikeShop struct {
	Id        primitive.ObjectID   `bson:"_id,omitempty"`
	LikedShop primitive.ObjectID `bson:"liked_shop,omitempty" json:"liked_shop,omitempty"`
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)

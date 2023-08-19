package shopmodel

import (
	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionName = "shop"

type ShopModel struct {
	common.SqlModel `bson:"inline"`
	Id              primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string               `json:"name,omitempty" bson:"name,omitempty"`
	Products        []primitive.ObjectID `json:"products,omitempty" bson:"products,omitempty"`
	LikedCount      int                  `json:"liked_count,omitempty" bson:"liked_count,omitempty"`
	UserId					primitive.ObjectID	 `json:"user_id,omitempty" bson:"user_id,omitempty"`
}

func GetCollectionName() string {
	return CollectionName
}

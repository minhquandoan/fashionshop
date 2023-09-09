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
	LikedCount      uint16               `json:"liked_count,omitempty" bson:"liked_count,omitempty"`
	UserId			primitive.ObjectID	 `json:"user_id,omitempty" bson:"user_id,omitempty"`
}

type ShopUpdateLikedCount struct {
	Id              primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Value			int16
}
// Constructor
func NewShopUpdateLikedCount(id primitive.ObjectID, value int16) *ShopUpdateLikedCount {
	return &ShopUpdateLikedCount{
		Id: id,
		Value: value,
	}
} 

// Get, Set for ShopUpdateLikedCount model
func (model *ShopUpdateLikedCount) GetId() primitive.ObjectID { return model.Id }
func (model *ShopUpdateLikedCount) GetValue() int16 { return model.Value }

func GetCollectionName() string {
	return CollectionName
}

package uploadmodel

import (
	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Upload struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	common.Image `bson:"inline"`
}
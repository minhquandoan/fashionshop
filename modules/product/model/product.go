package productmodel

import (
	"github.com/minhquandoan/fashionshop/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionName = "product"
)

type Product struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name *string `bson:"name,omitempty"`
	Description *string `bson:"description,omitempty"`
	Images []ProductImage `bson:"images,omitempty"`
	Properties bson.A `bson:"properties,omitempty"`
	Variants []Variant `bson:"variants,omitempty"`
	common.SqlModel `bson:"inline"`
}

type Variant struct {
	ErpID *string `bson:"erp_id,omitempty" json:"erp_id"`  
	StockedKU *string `bson:"sku,omitempty" json:"sku"`
	Price *string `bson:"price,omitempty"`
	Quantity *uint32 `bson:"quantity,omitempty"` 
	Image *common.Image `bson:"image,omitempty"`
}

type ProductProperties struct {
	Properties interface{} `bson:"properties,omitempty"`
}

type ProductImage struct {
	common.Image `bson:"inline"`
}

func GetCollectionName() string {
	return CollectionName
}
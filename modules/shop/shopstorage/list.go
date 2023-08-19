package shopstorage

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *ShopStorage) ListById(ctx context.Context, id *string) (*shopmodel.ShopModel, error) {
	coll := store.collection
	
	objID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, common.NewCustomError(err, "", "")
	}

	filter := bson.D{{"_id", objID}}
	matchStage := bson.D{{"$match", filter}}

	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if !cursor.Next(ctx) {
		return nil, common.ErrEntityNotFound("shop", mongo.ErrNilCursor)
	}

	var shop shopmodel.ShopModel
	if err := cursor.Decode(&shop); err != nil {
		return nil, common.ErrEntityNotFound("shop", mongo.ErrNilDocument)	
	}

	return &shop, nil
}

func (store *ShopStorage) ListShops(ctx context.Context, data *[]shopmodel.ShopModel) error {
	coll := store.collection

	filter := bson.D{{"status", 1}}
	matchStage := bson.D{{"$match", filter}}

	curs, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage})
	if err != nil {
		return common.ErrCannotListEntity("shop", err)
	}

	if err := curs.All(ctx, data); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}


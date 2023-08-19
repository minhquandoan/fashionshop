package userstorage

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (store *userStore) FindUser(context context.Context, filters *bson.D) (*usermodel.User, error) {
	coll := store.collection

	var user usermodel.User

	matchStage := bson.D{{"$match", filters}}

	cur, err := coll.Aggregate(context, mongo.Pipeline{matchStage})
	if err != nil {
		return nil, common.ErrDB(err)
	}

	defer cur.Close(context)

	if !cur.Next(context) {
		return nil, common.ErrEntityNotFound("user", mongo.ErrNilCursor)
	}

	if err := cur.Decode(&user); err != nil {
		return nil, common.ErrInternal(err)
	}

	return &user, nil
}

func (store *userStore) FindUserLikeShop(ctx context.Context, userId primitive.ObjectID) (*usermodel.UserLikeShopModel, error) {
	coll := store.collection

	curs, err := coll.Find(ctx, bson.M{"_id": userId})
	if err != nil {
		return nil, common.ErrCannotGetEntity("user", err)
	}

	defer curs.Close(ctx)

	if ok := curs.Next(ctx); !ok {
		return nil, common.ErrInternal(mongo.ErrNilCursor)
	}

	var result usermodel.UserLikeShopModel
	if err := curs.Decode(&result); err != nil {
		return nil, common.ErrInternal(mongo.ErrNilCursor)
	}

	return &result, nil
}

func (store *userStore) CheckUserLikeShop(ctx context.Context, filter *usermodel.UserLikeShop) (bool, error) {
	coll := store.collection

	count, err := coll.CountDocuments(ctx, bson.M{"_id": filter.Id, "liked_shop": filter.LikedShop})
	if err != nil {
		return false, common.ErrInternal(err)
	}

	if count != 1 {
		return false, nil
	}

	return true, nil
}
package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/component/tokenprovider/jwt"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/user/userbiz"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"github.com/minhquandoan/fashionshop/modules/user/userstorage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprint("wrong authen header"),
		fmt.Sprint("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(*appCtx.GetSecret())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		store := userstorage.NewUserStorage(db.GetCollection(appCtx.GetDbClient(), usermodel.CollectionName))
		biz := userbiz.NewSubUserLoginBiz(store)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		//ObjectID("645754870297851bb685e339") get the ID hex
		id, err := primitive.ObjectIDFromHex(payload.UserId[10:len(payload.UserId) - 2])
		if err != nil {
			panic(err)
		}
		filter := bson.D{{"_id", id}}
		user, err := biz.Store.FindUser(c, &filter)
		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set("user", user)
		c.Next()
	}
}
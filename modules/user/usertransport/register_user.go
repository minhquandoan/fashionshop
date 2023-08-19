package usertransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quan-doan/golang-mongo-example/component"
	"github.com/quan-doan/golang-mongo-example/component/hasher"
	"github.com/quan-doan/golang-mongo-example/db"
	"github.com/quan-doan/golang-mongo-example/modules/user/userbiz"
	"github.com/quan-doan/golang-mongo-example/modules/user/usermodel"
	"github.com/quan-doan/golang-mongo-example/modules/user/userstorage"
)

func RegisterUser(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user usermodel.User
		if err := ctx.ShouldBind(&user); err != nil {
			panic(err)
		}

		md5Hasher := hasher.NewMD5Hash()
		store := userstorage.NewUserStorage(db.GetCollection(appCtx.GetDbClient(), "user"))
		biz := userbiz.NewRegisterUserBiz(store, md5Hasher)

		id, err := biz.RegisterUser(ctx, &user)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": id,
		})
	}
}
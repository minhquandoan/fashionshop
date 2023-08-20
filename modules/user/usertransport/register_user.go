package usertransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/component/hasher"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/user/userbiz"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"github.com/minhquandoan/fashionshop/modules/user/userstorage"
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
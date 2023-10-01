package usertransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/component/hasher"
	"github.com/minhquandoan/fashionshop/component/tokenprovider/jwt"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/user/userbiz"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"github.com/minhquandoan/fashionshop/modules/user/userstorage"
)

func AccountLogin(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var account usermodel.UserAccount
		if err := c.ShouldBind(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": common.ErrInvalidRequest(err),
			})
			return
		}

		store := userstorage.NewUserStorage(db.GetCollection(appCtx.GetDbClient(), usermodel.CollectionName))
		tokenProvider := jwt.NewTokenJWTProvider(*appCtx.GetSecret())
		md5Hash := hasher.NewMD5Hash()

		biz := userbiz.NewUserLoginBiz(appCtx, store, tokenProvider, md5Hash, 60*60*24*7)
		accessToken, err := biz.Login(c, &account)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": accessToken,
		})
	}
}
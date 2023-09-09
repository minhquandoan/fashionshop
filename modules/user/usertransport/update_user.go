package usertransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/modules/shop/shopmodel"
	"github.com/minhquandoan/fashionshop/modules/shop/shopstorage"
	"github.com/minhquandoan/fashionshop/modules/user/userbiz"
	"github.com/minhquandoan/fashionshop/modules/user/usermodel"
	"github.com/minhquandoan/fashionshop/modules/user/userstorage"
)

func LikeShop(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userStore := userstorage.NewUserStorage(db.GetCollection(appCtx.GetDbClient(), usermodel.CollectionName))
		shopStore := shopstorage.NewShopStorage(db.GetCollection(appCtx.GetDbClient(), shopmodel.CollectionName))

		var data usermodel.UserLikeShop
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": common.ErrInvalidRequest(err),
			})
			return
		}

		biz := userbiz.NewUpdateUserBiz(userStore, shopStore, appCtx.GetPubSub())
		result, err := biz.UserLikeShop(c, &data)
		if err !=  nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": common.ErrCannotUpdateEntity("shop", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
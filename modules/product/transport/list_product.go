package producttransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProduct() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "Hello",
		})
	}
}
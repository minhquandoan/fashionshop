package uploadtransport

import (
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/component/uploadprovider"
	"github.com/minhquandoan/fashionshop/modules/upload/uploadbiz"
)

func Upload(ctx component.AppContext, uploadProvider *uploadprovider.Provider) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get File from request form
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}

		// Fetch location folder - default: /img
		folder := c.DefaultPostForm("folder", "img")
		
		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}

		defer file.Close()

		// Decode image
		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(err)
		}

		// fmt.Println(string(dataBytes))

		biz := uploadbiz.NewUploadBiz(*uploadProvider, nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, &folder, &fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": img,
		})
	}
}

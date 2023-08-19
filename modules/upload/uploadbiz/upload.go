package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/minhquandoan/fashionshop/common"
	"github.com/minhquandoan/fashionshop/component/uploadprovider"
)

type UploadStore interface {
	CreateImage(ctx context.Context, data *common.Image) (interface{}, error)
}

type uploadBiz struct {
	provider uploadprovider.Provider
	store UploadStore
}

func NewUploadBiz(provider uploadprovider.Provider, store UploadStore) *uploadBiz {
	return &uploadBiz{provider:provider, store: store}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName *string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)
	
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if strings.TrimSpace(*folder) == "" {
		*folder = "img"
	}

	fileExt := filepath.Ext(*fileName)
	*fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", *folder, *fileName))
	if err != nil {
		// log.Fatal("can not save file to cloud")
		return nil, common.ErrCannotCreateEntity("", err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		// log.Fatalln("Can not read image configuration")
		return 0, 0, common.ErrInvalidImage(err)
	}

	return img.Width, img.Height, nil
}
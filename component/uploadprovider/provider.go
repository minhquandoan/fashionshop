package uploadprovider

import (
	"context"

	"github.com/minhquandoan/fashionshop/common"
)

type Provider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

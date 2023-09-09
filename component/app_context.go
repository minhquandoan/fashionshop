package component

import (
	"github.com/minhquandoan/fashionshop/component/uploadprovider"
	"github.com/minhquandoan/fashionshop/pubsub"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetDbClient() *mongo.Client
	GetUploadProvider() *uploadprovider.Provider
	GetSecret() *string
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	dbClient *mongo.Client
	upload *uploadprovider.Provider
	secret *string
	pb pubsub.PubSub
}

func NewAppCtx(client *mongo.Client , provider *uploadprovider.Provider, secret *string, pb pubsub.PubSub) *appCtx {
	return &appCtx{dbClient: client, upload: provider, secret: secret, pb: pb}
}

func (ctx *appCtx) GetDbClient() *mongo.Client {
	return ctx.dbClient
}

func (ctx *appCtx) GetUploadProvider() *uploadprovider.Provider {
	return ctx.upload
}

func (ctx *appCtx) GetSecret() *string {
	return ctx.secret
}

func(ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.pb
}
package component

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetDbClient() *mongo.Client
	// GetUploadProvider() *uploadprovider.Provider
	// GetSecret() *string
}

type appCtx struct {
	dbClient *mongo.Client
	// upload *uploadprovider.Provider
	// secret *string
}

func NewAppCtx(client *mongo.Client /*, provider *uploadprovider.Provider, secret *string*/) *appCtx {
	return &appCtx{dbClient: client/*, upload: provider, secret: secret*/}
}

func (ctx *appCtx) GetDbClient() *mongo.Client {
	return ctx.dbClient
}

// func (ctx *appCtx) GetUploadProvider() *uploadprovider.Provider {
// 	return ctx.upload
// }

// func (ctx *appCtx) GetSecret() *string {
// 	return ctx.secret
// }
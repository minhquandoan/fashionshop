package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/component/uploadprovider"
	"github.com/minhquandoan/fashionshop/db"
	"github.com/minhquandoan/fashionshop/middleware"
	producttransport "github.com/minhquandoan/fashionshop/modules/product/transport"
	"github.com/minhquandoan/fashionshop/modules/shop/shoptransport"
	"github.com/minhquandoan/fashionshop/modules/upload/uploadtransport"
	"github.com/minhquandoan/fashionshop/modules/user/usertransport"
	"github.com/minhquandoan/fashionshop/pubsub/localpb"
	"github.com/minhquandoan/fashionshop/subscriber"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	fmt.Println("Starting the Application ...")

	// Application shoud be started in 10 secs
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Open DB
	clientDb, err := db.ConnectDb(ctx, os.Getenv("databasePath"))
	if err != nil {
		panic(err)
	}
	log.Println("Accessed to Database Successfully!!")
	
	//Close DB
	defer func ()  {
		if err := clientDb.Disconnect(ctx); err != nil {
			panic(err)
		}

		fmt.Println("Closing the Application ...")
	}()

	//S3 Service Config
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	// System Secrets
	secret := os.Getenv("SYSTEM_SECRET")

	// Run all APIs
	runService(clientDb, s3Provider, &secret)
}

func runService(clientDb *mongo.Client, uploadProvider uploadprovider.Provider, passSecret *string) error {
	appCtx := component.NewAppCtx(clientDb, &uploadProvider, passSecret, localpb.NewLocalPubSub())

	r := gin.Default()

	// PubSub mechanism
	subscriber.SetUp(appCtx)

	// Middelwares
	r.Use(middleware.Recover(appCtx))

	//API versions
	v1 := r.Group("/v1")
	
	// Products API (GET, POST, UPDATE, DELETE)
	productGr := v1.Group("/product", middleware.RequiredAuth(appCtx)) 
	{
		productGr.GET("/get", producttransport.ListProduct(appCtx))
		productGr.POST("/getby", producttransport.ListProductsByFilters(appCtx))
		productGr.POST("/create", producttransport.CreateOneProduct(appCtx))
		productGr.PATCH("/updatebyid/:id", producttransport.UpdateOneProduct(appCtx))
	}

	// Upload Image API
	uploadGroup := v1.Group("/upload", middleware.RequiredAuth(appCtx))
	{
		uploadGroup.POST("/", uploadtransport.Upload(appCtx, &uploadProvider))
	}

	// User APIs
	userGroup := v1.Group("/user")
	{
		userGroup.POST("/register", usertransport.RegisterUser(appCtx))
		userGroup.POST("/login", usertransport.AccountLogin(appCtx))
		userGroup.POST("/likeshop", middleware.RequiredAuth(appCtx), usertransport.LikeShop(appCtx))
	}

	// Shop APIs
	shopGroup := v1.Group("/shop", middleware.RequiredAuth(appCtx)) 
	{
		shopGroup.POST("/add", shoptransport.AddShop(appCtx))
		shopGroup.GET("/get:id", shoptransport.ListShopById(appCtx))
		shopGroup.PATCH("/incrlike/:id", shoptransport.IncreaseLikedCount(appCtx))
	}

	return r.Run()
}
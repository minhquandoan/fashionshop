package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minhquandoan/fashionshop/component"
	"github.com/minhquandoan/fashionshop/db"
	producttransport "github.com/minhquandoan/fashionshop/modules/product/transport"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	fmt.Println("Starting the Application ...")

	// Application shoud be started in 10 secs
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Open DB
	clientDb, _ := db.ConnectDb(ctx, os.Getenv("databasePath"))
	log.Println("Accessed to Database Successfully!!")
	
	//Close DB
	defer func ()  {
		if err := clientDb.Disconnect(ctx); err != nil {
			panic(err)
		}

		fmt.Println("Closing the Application ...")
	}()

	// Run all APIs
	runService(clientDb)
}

func runService(clientDb *mongo.Client) error {
	appCtx := component.NewAppCtx(clientDb)

	r := gin.Default()

	//API 
	
	// Products API (GET, POST, UPDATE, DELETE)
	productGr := r.Group("/v1/product") 
	{
		productGr.GET("/get", producttransport.ListProduct(appCtx))
	}

	return r.Run()
}
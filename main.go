package main

import (
	"context"
	"log"
	"sgs_fullstack_challenge/configs"
	"sgs_fullstack_challenge/middlewares"
	"sgs_fullstack_challenge/routes"
	"sgs_fullstack_challenge/seeds"
	"sgs_fullstack_challenge/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Get configs
	config := configs.LoadConfig()

	// We Connect and initialize the DB so we can Inject it on other functions
	db, err := utils.ConnectToDB(config)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Client().Disconnect(context.Background())

	productCollection := db.Collection(config.MongoCollectionName)
	//Check if we have to run a seeder
	if utils.IsCollectionEmpty(productCollection) {
		seeds.ProductSeeder(productCollection, 1000)
	}
	//Middlewares
	router.Use(cors.Default())
	router.Use(middlewares.DatabaseMiddleware(db))
	//Routes
	routes.SetupRoutes(router, *productCollection)

	router.Run("localhost:8000")
}

package routes

import (
	"sgs_fullstack_challenge/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, collection mongo.Collection) {
	productCtrl := &controllers.ProductController{}

	router.GET("/products", productCtrl.GetAllProducts)
}

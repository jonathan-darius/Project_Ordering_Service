package main

import (
	"WebService/App/routes"
	"WebService/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/logrusorgru/aurora/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title           Gin Mini Product Ordering Service
// @version         1.0
// @description     A Mini Product Ordering Service API.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Token
var (
	router *gin.Engine
	port   string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router = gin.Default()
	//router.Use(gin.Logger())
}

func main() {
	version := router.Group("/api/v1")
	routes.ProductRoutes(version)
	routes.AuthRoutes(version)
	routes.UserRoutes(version)
	routes.ProductManagement(version)
	routes.CartManagement(version)
	routes.PurchaseManagement(version)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"
	docsUrl := fmt.Sprintf("http://localhost:%s/swagger/index.html", port)
	fmt.Println("\nSwagger Docs Url: ", aurora.Green(docsUrl).Hyperlink(
		docsUrl,
	))

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

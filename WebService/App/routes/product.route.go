package routes

import (
	"WebService/App/controllers"
	"WebService/App/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(req *gin.RouterGroup) {
	group := req.Group("Product")
	group.POST("/Search", controllers.SearchProduct())
	group.GET("/:pid", controllers.GetProduct())
}

func ProductManagement(req *gin.RouterGroup) {
	req.Use(middleware.Authenticate())
	group := req.Group("Product")

	group.POST("/AddProduct", controllers.AddProduct())
	group.PATCH("/UpdateProduct", controllers.UpdateProduct())
	group.DELETE("/Delete/:pid", controllers.DeleteProduct())

	group.POST("/AddCategory", controllers.AddCategory())
	group.POST("/RemoveCategory", controllers.RemoveCategory())

	group.POST("/AddStock", controllers.AddProductStock())

	group.POST("/AddImage", controllers.AddProductImage())
}

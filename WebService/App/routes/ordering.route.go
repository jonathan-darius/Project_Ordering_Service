package routes

import (
	"WebService/App/controllers"
	"WebService/App/middleware"
	"github.com/gin-gonic/gin"
)

func CartManagement(req *gin.RouterGroup) {
	req.Use(middleware.Authenticate())
	group := req.Group("Cart")

	group.POST("/AddItem", controllers.AddCart())
	group.DELETE("/RemoveItem/:pID", controllers.RemoveCartItem())

	group.DELETE("/EmptyCart", controllers.EmptyCart())

	group.GET("/:uID", controllers.GetCart())
}

func PurchaseManagement(req *gin.RouterGroup) {
	req.Use(middleware.Authenticate())
	group := req.Group("Transaction")

	group.POST("/Purchase/:uID", controllers.PurchaseCart())

	group.GET("/:tID", controllers.GetTransaction())
	group.GET("/All", controllers.GetAllTransactions())

	group.POST("/Rating", controllers.AddRating())
}

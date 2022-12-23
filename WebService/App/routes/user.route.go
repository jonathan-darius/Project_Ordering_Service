package routes

import (
	"WebService/App/controllers"
	"WebService/App/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(req *gin.RouterGroup) {
	group := req.Group("Auth")
	group.POST("/register", controllers.RegisterController())
	group.POST("/login", controllers.LoginController())
}

func UserRoutes(req *gin.RouterGroup) {
	req.Use(middleware.Authenticate())
	group := req.Group("User")
	group.GET("/verification/:email", controllers.SendAccountVerification())
	group.POST("/verification", controllers.AccountVerification())
	group.GET("/profile", controllers.ShowUser())
	group.GET("/profile/all", controllers.ListUser())
	group.DELETE("/delete/:uid", controllers.DeleteUser())
	group.PATCH("/update", controllers.UpdateUser())
	group.POST("/image", controllers.UploadImage())
}

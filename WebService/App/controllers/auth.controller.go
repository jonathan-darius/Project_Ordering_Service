package controllers

import (
	gRPCFunc "WebService/App/gRPC_Configs/User"
	"WebService/App/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
)

var SafeGcp = true

// LoginController 	godoc
// @Summary      	Login Form
// @Description  	Login Form
// @Tags         	Auth
// @Accept 			multipart/form-data
// @Produce 		json
// @Param 			email 		formData 	string 	true "email"
// @Param 			password 	formData 	string 	true "password"
// @Router			/Auth/login [post]
func LoginController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginData models.FormLoginData
		if err := ctx.Bind(&loginData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "Failed To Bind Form",
			})
			return
		}

		res, err := gRPCFunc.UserLogin(loginData)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": status.Convert(err).Message(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
			"Data":    res,
		})
	}
}

// RegisterController godoc
// @Summary      	Register Form
// @Description  	Register Form
// @Tags         	Auth
// @Accept 			json
// @Produce 		json
// @Param       	user body		models.RegisterData true "Register User"
// @Router       	/Auth/register [post]
func RegisterController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if SafeGcp {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden",
			})
			return
		}
		var register models.RegisterData
		if err := ctx.BindJSON(&register); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		register.Role = "user"
		res, err := gRPCFunc.RegisterUser(register)

		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": status.Convert(err).Message(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
			"Data":    res,
		})
	}
}

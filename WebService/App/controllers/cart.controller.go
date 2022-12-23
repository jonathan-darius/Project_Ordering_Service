package controllers

import (
	gRPCFunc "WebService/App/gRPC_Configs/Ordering"
	"WebService/App/models"
	"WebService/App/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
)

// AddCart godoc
// @Summary      	Add Product to Cart
// @Description  	Add Product to Cart
// @Tags         	Cart
// @Accept 			json
// @Produce 		json
// @Param       	product body	models.CartItemSend true "Add Product to Cart "
// @Router       	/Cart/AddItem [post]
// @Security 		ApiKeyAuth
func AddCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
			})
			return
		}

		var newItem models.CartItemSend
		if err := ctx.BindJSON(&newItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		newItem.UserID = userID

		if err := gRPCFunc.AddCart(newItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": status.Convert(err).Message(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
		})

	}
}

// RemoveCartItem 	godoc
// @Summary      	Remove Product from Cart
// @Description  	Remove Product from Cart
// @Tags         	Cart
// @Accept 			json
// @Produce 		json
// @Param       	pID 	path 	string 		true 	"Remove Product from Cart "
// @Router       	/Cart/RemoveItem/{pID} [delete]
// @Security 		ApiKeyAuth
func RemoveCartItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
			})
			return
		}

		if err := gRPCFunc.RemoveCartItem(models.CartItemSend{
			ProductID: ctx.Param("pID"),
			UserID:    userID,
		}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": status.Convert(err).Message(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
		})
	}
}

// EmptyCart	 	godoc
// @Summary      	Remove All Product from Cart
// @Description  	Remove All Product from Cart
// @Tags         	Cart
// @Accept 			json
// @Produce 		json
// @Param       	uID 	path 	string 		true 	"Remove All Product from Cart "
// @Router       	/Cart/EmptyCart/{uID} [delete]
// @Security 		ApiKeyAuth
func EmptyCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
			})
			return
		}
		if err := gRPCFunc.EmptyUserCart(userID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": status.Convert(err).Message(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
		})
	}
}

// GetCart		 	godoc
// @Summary      	Get All Product from Cart
// @Description  	Get All Product from Cart
// @Tags         	Cart
// @Accept 			json
// @Produce 		json
// @Param       	uID 	path 	string 		true 	"Get All Product from Cart "
// @Router       	/Cart/{uID} [get]
// @Security 		ApiKeyAuth
func GetCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		isAdmin := utils.CheckAdmin(ctx.GetString("role"))

		userParam := ctx.Param("uID")
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
				"Data":    "",
			})
			return
		}
		if !isAdmin && (userParam != userID) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Access",
				"Data":    "",
			})
			return
		}
		res, err := gRPCFunc.GetUserCart(userParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": status.Convert(err).Message(),
				"Data":    "",
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

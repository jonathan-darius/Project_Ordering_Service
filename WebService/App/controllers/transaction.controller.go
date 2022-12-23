package controllers

import (
	gRPCFunc "WebService/App/gRPC_Configs/Ordering"
	"WebService/App/models"
	"WebService/App/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

// PurchaseCart		 	godoc
// @Summary      	Purchase All Product from Cart
// @Description  	Purchase All Product from Cart
// @Tags         	Transaction
// @Accept 			json
// @Produce 		json
// @Param       	uID 	path 	string 		true 	"Purchase All Product from Cart "
// @Router       	/Transaction/Purchase/{uID} [post]
// @Security 		ApiKeyAuth
func PurchaseCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		isAdmin := utils.CheckAdmin(ctx.GetString("role"))
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
			})
			return
		}
		userParam := ctx.Param("uID")
		if !isAdmin && (userParam != userID) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Access",
				"Data":    "",
			})
			return
		}
		res, err := gRPCFunc.PurchaseCart(userParam)
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

// GetTransaction	godoc
// @Summary      	Detai Transaction
// @Description  	Detai Transaction
// @Tags         	Transaction
// @Accept 			json
// @Produce 		json
// @Param       	tID 	path 	string 		true 	"Detai Transaction"
// @Router       	/Transaction/{tID} [get]
// @Security 		ApiKeyAuth
func GetTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		isAdmin := utils.CheckAdmin(ctx.GetString("role"))
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
			})
			return
		}

		tID := ctx.Param("tID")
		res, err := gRPCFunc.DetailTransaction(tID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": status.Convert(err).Message(),
			})
		}

		if !isAdmin && (res.GetUserId() != userID) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Access",
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

// AddRating		godoc
// @Summary      	Add Product Rating
// @Description  	Add Product Rating
// @Tags         	Transaction
// @Accept 			json
// @Produce 		json
// @Param       	rating 	body 	models.ProductRating 	true 	"Add Product Rating"
// @Router       	/Transaction/Rating [post]
// @Security 		ApiKeyAuth
func AddRating() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
			})
			return
		}

		var newRating *models.ProductRating
		if err := ctx.BindJSON(&newRating); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		newRating.UserID = userID
		if err := gRPCFunc.AddRating(newRating); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
		})
	}
}

// GetAllTransactions godoc
// @Summary      	Show  All User Transactions
// @Description  	Show  All User Transactions
// @Tags         	Transaction
// @Produce 		json
// @Param 			page    		query	int  false  "page number"  		Format(number)
// @Param 			recordPerPage   query   int  false  "data per page"  	Format(number)
// @Router       	/Transaction/All [get]
// @Security 		ApiKeyAuth
func GetAllTransactions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("uid")
		if !utils.CheckVerifyAccount(userID) {
			ctx.JSON(http.StatusLocked, gin.H{
				"Status":  http.StatusLocked,
				"Message": "Unverified User",
			})
			return
		}

		recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 5
		}
		page, err1 := strconv.Atoi(ctx.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		res, err := gRPCFunc.GetAllUserTransaction(userID, int32(page), int32(recordPerPage))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": status.Convert(err).Message(),
			})
		}
		data := utils.Pagination{
			Page:   int32(page),
			Record: int32(recordPerPage),
			Data:   res,
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
			"Data":    data,
		})
	}
}

package controllers

import (
	gRPCFunc "WebService/App/gRPC_Configs/Product"
	"WebService/App/models"
	"WebService/App/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
)

// AddProduct godoc
// @Summary      	Add Product
// @Description  	Add Product
// @Tags         	Product Management
// @Accept 			json
// @Produce 		json
// @Param       	product body	models.Product true "Add Product "
// @Router       	/Product/AddProduct [post]
// @Security 		ApiKeyAuth
func AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !utils.CheckAdmin(ctx.GetString("role")) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		var registerProduct models.Product
		if err := ctx.BindJSON(&registerProduct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		registerProduct.CreatedBy = ctx.GetString("uid")

		res, err := gRPCFunc.AddProduct(registerProduct)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
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

// UpdateProduct 	godoc
// @Summary      	Add Product
// @Description  	Add Product
// @Tags         	Product Management
// @Accept 			json
// @Produce 		json
// @Param       	product body	models.UpdateProduct true "Update Product"
// @Router       	/Product/UpdateProduct [patch]
// @Security 		ApiKeyAuth
func UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateProduct *models.UpdateProduct
		if err := ctx.BindJSON(&updateProduct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !utils.CheckAdmin(ctx.GetString("role")) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		err := gRPCFunc.ProductUpdate(updateProduct)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": status.Convert(err).Message(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
			"Data":    "",
		})

	}
}

// DeleteProduct 	godoc
// @Summary      	Delete Product
// @Description  	Delete Product
// @Tags         	Product Management
// @Accept	       	json
// @Produce      	json
// @Param		 	pid  	path 	string 	true  	"Delete Product by ID"
// @Router       	/Product/Delete/{pid} [delete]
// @Security 		ApiKeyAuth
func DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !utils.CheckAdmin(ctx.GetString("role")) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		pid := ctx.Param("pid")
		err := gRPCFunc.DelProduct(pid)
		if err != nil {
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

// AddCategory 	godoc
// @Summary      	Add Category Product
// @Description  	Add Category Product
// @Tags         	Product Management
// @Accept	       	json
// @Produce      	json
// @Param       	product body	models.ProductArr true "Add Category Product"
// @Router       	/Product/AddCategory [post]
// @Security 		ApiKeyAuth
func AddCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !utils.CheckAdmin(ctx.GetString("role")) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		var categoryProduct *models.ProductArr
		if err := ctx.BindJSON(&categoryProduct); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Status":  http.StatusInternalServerError,
				"Message": status.Convert(err).Message(),
			})
			return
		}
		if err := gRPCFunc.AddCategory(categoryProduct); err != nil {
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

// RemoveCategory 	godoc
// @Summary      	Remove Category Product
// @Description  	Remove Category Product
// @Tags         	Product Management
// @Accept	       	json
// @Produce      	json
// @Param       	product body	models.ProductArr true "Remove Category Product"
// @Router       	/Product/RemoveCategory [post]
// @Security 		ApiKeyAuth
func RemoveCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !utils.CheckAdmin(ctx.GetString("role")) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		var categoryProduct *models.ProductArr
		if err := ctx.BindJSON(&categoryProduct); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Status":  http.StatusInternalServerError,
				"Message": status.Convert(err).Message(),
			})
			return
		}
		if err := gRPCFunc.RemoveCategory(categoryProduct); err != nil {
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

// AddProductStock 	godoc
// @Summary      	Add Stock Product
// @Description  	Add Stock Product
// @Tags         	Product Management
// @Accept	       	json
// @Produce      	json
// @Param       	product body	models.ProductStock true "Add Stock Product"
// @Router       	/Product/AddStock [post]
// @Security 		ApiKeyAuth
func AddProductStock() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !utils.CheckAdmin(ctx.GetString("role")) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		var stockProduct *models.ProductStock
		if err := ctx.BindJSON(&stockProduct); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Status":  http.StatusInternalServerError,
				"Message": status.Convert(err).Message(),
			})
			return
		}
		if err := gRPCFunc.AddStock(stockProduct); err != nil {
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

// AddProductImage 	godoc
// @Summary      	Product Image Form
// @Description  	Product Image Form
// @Tags         	Product Management
// @Accept 			multipart/form-data
// @Param 	id 		formData string true "id"
// @Param 	image 	formData file 	true "image"
// @Success      	200  {string}	{}
// @Router       	/Product/AddImage [post]
// @Security 		ApiKeyAuth
func AddProductImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, _ := ctx.FormFile("image")
		uid, _ := ctx.GetPostForm("id")
		err := gRPCFunc.ProductImage(file, uid)
		if err != nil {
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

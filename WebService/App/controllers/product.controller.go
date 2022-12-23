package controllers

import (
	gRPCFunc "WebService/App/gRPC_Configs/Product"
	"WebService/App/models"
	"WebService/App/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

// SearchProduct godoc
// @Summary      	Search Product
// @Description  	Search Product
// @Tags         	Product
// @Accept 			json
// @Produce 		json
// @Param 			page    		query	int  false  "page number"  		Format(number)
// @Param 			recordPerPage   query   int  false  "data per page"  	Format(number)
// @Param       	product body	models.SearchProduct false "Search Product "
// @Router       	/Product/Search [post]
func SearchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 5
		}
		page, err1 := strconv.Atoi(ctx.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		var searchProduct *models.SearchProduct
		if err := ctx.BindJSON(&searchProduct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := gRPCFunc.SearchProduct(searchProduct, int32(page), int32(recordPerPage))
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

// GetProduct godoc
// @Summary      Detail Product
// @Description  Detail Product
// @Tags         Product
// @Produce      json
// @Param		 pid  path string true  "Get Detail Product"
// @Router       /Product/{pid} [get]
func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pID := ctx.Param("pid")
		res, err := gRPCFunc.GetProduct(pID)
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
			"Data": models.Product{
				ID:        res.Id,
				Name:      res.Name,
				Price:     res.Price,
				Stock:     res.Stock,
				Sold:      res.Sold,
				Rating:    res.Rating,
				Rated:     res.Rated,
				Desc:      res.Desc,
				Category:  res.Category,
				Image:     res.Image,
				CreatedBy: res.CreatedBy,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.CreatedAt,
			},
		})
	}
}

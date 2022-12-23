package controllers

import (
	gRPCFunc "WebService/App/gRPC_Configs/User"
	"WebService/App/models"
	"WebService/App/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

// SendAccountVerification godoc
// @Summary      Account Verification User
// @Description  Account Verification User
// @Tags         User
// @Produce      json
// @Param		 email  path string true  "Account Verification User by Email"
// @Router       /User/verification/{email} [get]
// @Security 	ApiKeyAuth
func SendAccountVerification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mail := ctx.Param("email")

		isAdmin := utils.CheckAdmin(ctx.GetString("role"))
		if !isAdmin {
			userCheck, err := utils.CheckAuthUser(mail, ctx.GetString("uid"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Status":  http.StatusBadRequest,
					"Message": status.Convert(err).Message(),
				})
				return
			}
			if !userCheck {
				ctx.JSON(http.StatusForbidden, gin.H{
					"Status":  http.StatusForbidden,
					"Message": "Forbidden Request: Login Information & Email Request Not Same",
				})
				return
			}
		}

		if err := gRPCFunc.GetAccountVerification(mail); err != nil {
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

// AccountVerification godoc
// @Summary      	Account Verification Form
// @Description  	Account Verification Form
// @Tags         	User
// @Accept 			json
// @Produce 		json
// @Param       	user body		models.AccountVerificationData true "Account Verification User"
// @Router       	/User/verification [post]
// @Security 		ApiKeyAuth
func AccountVerification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			verifyData models.AccountVerificationData
			mail       string
		)

		if err := ctx.BindJSON(&verifyData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isAdmin := utils.CheckAdmin(ctx.GetString("role"))
		if !isAdmin || verifyData.Email == "" {
			raw, err := gRPCFunc.SearchUserByID(ctx.GetString("uid"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Status":  http.StatusBadRequest,
					"Message": status.Convert(err).Message(),
				})
				return
			}
			mail = raw.Email
		} else {
			mail = verifyData.Email
		}
		err := gRPCFunc.AccountVerification(mail, verifyData.Token)
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

// ShowUser godoc
// @Summary      	Show User Form
// @Description  	Show User Form
// @Tags         	User
// @Produce 		json
// @Param		 	uid  query string false  "Show User by ID"
// @Router       	/User/profile [get]
// @Security 		ApiKeyAuth
func ShowUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Query("uid")
		if userID == "" {
			userID = ctx.GetString("uid")
		}
		res, err := gRPCFunc.SearchUserByID(userID)
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
			"Data":    utils.ResponseUser(res),
		})
	}
}

// ListUser godoc
// @Summary      	Show User
// @Description  	Show User
// @Tags         	User
// @Produce 		json
// @Param 			page    		query	int  false  "page number"  		Format(number)
// @Param 			recordPerPage   query   int  false  "data per page"  	Format(number)
// @Router       	/User/profile/all [get]
// @Security 		ApiKeyAuth
func ListUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isAdmin := utils.CheckAdmin(ctx.GetString("role"))
		if !isAdmin {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request: You Have No Access",
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
		res, err := gRPCFunc.ListUser(int32(page), int32(recordPerPage))
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

// DeleteUser godoc
// @Summary      Delete Account
// @Description  Delete Account
// @Tags         User
// @Produce      json
// @Param		 uid  path string true  "Delete Account by Email"
// @Router       /User/delete/{uid} [delete]
// @Security 	ApiKeyAuth
func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid := ctx.Param("uid")
		isAdmin := utils.CheckAdmin(ctx.GetString("role"))
		if !isAdmin && (ctx.GetString("uid") != uid) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		err := gRPCFunc.DelUser(uid)
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

// UpdateUser godoc
// @Summary      	Update Account
// @Description  	Update Account
// @Tags        	User
// @Accept 			json
// @Produce 		json
// @Param       	user body		models.UpdateModel true "Update User"
// @Router       	/User/update [patch]
// @Security 		ApiKeyAuth
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *models.UpdateModel
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "Failed To Bind Form",
			})
			return
		}
		isAdmin := utils.CheckAdmin(ctx.GetString("role"))
		if !isAdmin && (ctx.GetString("uid") != user.ID) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Status":  http.StatusForbidden,
				"Message": "Forbidden Request",
			})
			return
		}
		err := gRPCFunc.UpdateUser(user)
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

// UploadImage godoc
// @Summary      Image Form
// @Description  Post Form
// @Tags         User
// @Accept multipart/form-data
// @Param id formData string true "id"
// @Param image formData file true "image"
// @Success      200  {string}	{}
// @Router       /User/image [post]
// @Security 		ApiKeyAuth
func UploadImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, _ := ctx.FormFile("image")
		uid, _ := ctx.GetPostForm("id")
		err := gRPCFunc.SendImage(file, uid)
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

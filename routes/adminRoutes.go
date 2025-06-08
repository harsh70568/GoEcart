package routes

import (
	"goEcart/controllers"
	"goEcart/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/api/admin/v1")
	{
		/* Admin routes */
		admin.POST("/login", controllers.AdminLogin())
		admin.POST("/signup", controllers.AdminSignup())
		admin.POST("/logout", middleware.AdminAuth(), controllers.AdminLogout())

		/* Brand management routes */
		admin.PUT("/brand/editbrand/:id", middleware.AdminAuth(), controllers.EditBrand())
		admin.GET("/brand", middleware.AdminAuth(), controllers.GetBrand())

		/* User management routes */
		admin.GET("/user/viewuser", middleware.AdminAuth(), controllers.ViewAllUser())
		admin.GET("/user/searchuser", middleware.AdminAuth(), controllers.AdminSearchUser())
		admin.PUT("user/blockusers", middleware.AdminAuth(), controllers.AdminBlockUser())
		admin.PUT("user/edituserprofile/:id", middleware.AdminAuth(), controllers.AdminEditUserProfile())

		/* Product management routes */
		admin.POST("/addbrand", middleware.AdminAuth(), controllers.AddBrand())
		admin.POST("/addcategories", middleware.AdminAuth(), controllers.AddCategories())
		admin.POST("addproduct", middleware.AdminAuth(), controllers.AddProduct())

		/* Coupon routes */
		admin.POST("/coupon/add", middleware.AdminAuth(), controllers.AddCoupon())
		admin.POST("/coupon/validatecoupon", middleware.AdminAuth(), controllers.ValidateCoupon())
	}
}

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
		// admin.GET("/profile", middleware.AdminAuth(), controllers.AdminProfile())
		// admin.GET("/adminvalidate", middleware.AdminAuth(), controllers.ValidateAdmin())

		/* Brand management routes */
		admin.PUT("/brand/editbrand/:id", middleware.AdminAuth(), controllers.EditBrand())
		admin.GET("/brand", middleware.AdminAuth(), controllers.GetBrand())

		/* User management routes */
		admin.GET("/user/viewuser", middleware.AdminAuth(), controllers.ViewAllUser())
		admin.GET("/user/searchuser", middleware.AdminAuth(), controllers.AdminSearchUser())
		admin.PUT("user/blockusers", middleware.AdminAuth(), controllers.AdminBlockUser())
		admin.PUT("user/edituserprofile/:id", middleware.AdminAuth(), controllers.AdminEditUserProfile())

		/* Coupon routes */
		admin.POST("/coupon/add", middleware.AdminAuth(), controllers.AddCoupon())
		admin.POST("/coupon/validatecoupon", middleware.AdminAuth(), controllers.ValidateCoupon())
	}
}

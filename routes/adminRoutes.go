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

	}
}

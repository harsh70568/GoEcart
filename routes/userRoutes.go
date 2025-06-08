package routes

import (
	"goEcart/controllers"
	"goEcart/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/api/user/v1")
	{
		/* User routes */
		user.POST("/signup", controllers.UserSingup())
		user.POST("/signup/otpvalidate", controllers.OTPValidate())
		user.POST("/login", controllers.UserLogin())

		/* User profile management routes */
		user.GET("/viewprofile/:id", middleware.UserAuth(), controllers.ViewUserProfile())
		user.PUT("/editaddress/:id", middleware.UserAuth(), controllers.EditUserAddress())
		user.GET("searchaddress/:id", middleware.UserAuth(), controllers.SearchAddress())
	}
}

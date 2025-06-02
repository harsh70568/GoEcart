package routes

import (
	"goEcart/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/api/user/v1")
	{
		/* User routes */
		user.POST("/signup", controllers.UserSingup())
		user.POST("/signup/otpvalidate", controllers.OTPValidate())
		user.POST("/login", controllers.UserLogin())
	}
}

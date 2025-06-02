package controllers

import (
	"goEcart/db"
	"goEcart/models"
	"goEcart/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UserSingup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Check if already exists */
		var existingUser models.User
		if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user already exists"})
			return
		}

		/* Hash the password */
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in hashing the password"})
			return
		}
		user.Password = hashedPassword

		/* Send an email for OTP */
		otp := VerifyOTP(user.Email)
		user.OTP = otp

		/* Save to db */
		if err := db.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in saving to the database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Signup succesfull",
		})
	}
}

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Check if the email exists */
		var existingUser models.User
		if err := db.DB.Where("email = ?", user.Email).First(&existingUser); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email doesn't exists"})
			return
		}

		/* Check if user is blocked or not */
		if existingUser.IsBlocked {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User is blocked"})
			return
		}

		/* Match the passwords */
		if err := utils.VerifyPassword(existingUser.Password, user.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Password do not match"})
			return
		}

		/* Generate Token */
		token, err := utils.GenerateToken(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in generating token"})
			return
		}

		/* Set the cookies */
		c.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "Login succesfull",
			"id":      user.ID,
			"token":   token,
		})
	}
}

package controllers

import (
	"goEcart/db"
	"goEcart/models"
	"goEcart/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminSignup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var admin models.Admin
		if err := c.ShouldBindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Check if email already exists */
		var existingAdmin models.Admin
		if err := db.DB.Where("email = ?", admin.Email).Find(&existingAdmin).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}

		/* Hash the password */
		password, err := utils.HashPassword(admin.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in hashing the password"})
			return
		}
		admin.Password = password

		/* Save user in database */
		if err := db.DB.Create(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in saving to database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Signup succesfull"})
	}
}

func AdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var admin models.Admin
		if err := c.ShouldBindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		if admin.Email == "" || admin.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing crendetials"})
			return
		}

		/* Check if email does exists */
		var existingAdmin models.Admin
		if err := db.DB.Where("email = ?", admin.Email).First(&existingAdmin).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email already exists"})
			return
		}

		/* Verify the password */
		if err := utils.VerifyPassword(existingAdmin.Password, admin.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"erorr": "Password do not match"})
			return
		}

		/* Generate Token */
		token, err := utils.GenerateToken(admin.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in generating token"})
			return
		}

		/* Set Cookie */
		c.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "login sucessfull",
			"id":      admin.ID,
			"token":   token,
		})
	}
}

func AdminLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Clear the cookie */
		c.SetCookie("token", "", -1, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "Logout sucesfully"})
	}
}

package controllers

import (
	"goEcart/db"
	"goEcart/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ViewUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Extract the user id from the request */
		Sid := c.Param("id")
		id, err := strconv.Atoi(Sid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in getting user id"})
			return
		}

		/* Check if there is any user with same user id */
		var exisitngUser models.User
		if err := db.DB.Where("id = ?", id).First(&exisitngUser).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no user with the requested user id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User found",
			"details": exisitngUser,
		})
	}
}

func EditUserAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in getting id"})
			return
		}

		var exisitngAddress models.Address
		if err := c.ShouldBindJSON(&exisitngAddress).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}
		exisitngAddress.UserID = uint(id)

		/* Update where the user ID matches */
		if err := db.DB.Model(&models.Address{}).Where("user_id = ?", id).Updates(exisitngAddress).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no address with the requestd user id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":         "address updated succesfully",
			"updated_address": exisitngAddress,
		})
	}
}

func SearchAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in getting id"})
			return
		}

		/* Check if there is any address with the requested userID */
		var existingAddress models.Address
		if err := db.DB.Where("user_id = ?", id).First(&existingAddress).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no address found with requested user id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "address found",
			"data":    existingAddress,
		})
	}
}

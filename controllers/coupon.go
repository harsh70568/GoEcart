package controllers

import (
	"fmt"
	"goEcart/db"
	"goEcart/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Accept input of coupon and check if coupon with
same code exists in db or not
*/
func AddCoupon() gin.HandlerFunc {
	return func(c *gin.Context) {
		type CouponInput struct {
			CouponCode    string  `json:"coupon_code"`
			Year          uint    `json:"year"`
			Month         uint    `json:"month"`
			Day           uint    `json:"day"`
			DiscountPrice float64 `json:"discount_price"`
		}
		var coupon CouponInput
		if err := c.ShouldBindJSON(&coupon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Create expiry date */
		expiryDate := time.Date(int(coupon.Year), time.Month(coupon.Month), int(coupon.Day), 0, 0, 0, 0, time.UTC)

		/* Check if coupon already exists */
		var existingCoupon models.Coupon
		if err := db.DB.Where("coupon_code = ?", coupon.CouponCode).First(&existingCoupon).Error; err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Coupon code already taken"})
			return
		}

		/* Create coupon */
		newCoupon := models.Coupon{
			CouponCode:    coupon.CouponCode,
			DiscountPrice: coupon.DiscountPrice,
			Expired:       expiryDate,
		}

		if err := db.DB.Create(&newCoupon).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in saving new coupon"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "coupon added succesfully",
			"id":      newCoupon.ID,
		})
	}
}

/* Check if same coupon code exists and is not expired */
func ValidateCoupon() gin.HandlerFunc {
	return func(c *gin.Context) {
		type CouponInput struct {
			Coupon string
		}
		var coupon CouponInput
		if err := c.ShouldBindJSON(&coupon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Check if exists in db or not */
		var existingCoupon models.Coupon
		if err := db.DB.Where("coupon_code = ?", coupon.Coupon).First(&existingCoupon).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Coupon code %s does not exists", coupon.Coupon)})
			return
		}

		/* @TODO - Check expiry */
		c.JSON(http.StatusOK, gin.H{"message": "coupon is valid"})
	}
}

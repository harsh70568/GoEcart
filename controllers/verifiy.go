package controllers

import (
	"crypto/rand"
	"fmt"
	"goEcart/db"
	"goEcart/models"
	"math/big"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VerifyOTP(email string) string {
	OTP, err := getRandomNum()
	if err != nil {
		return ""
	}

	sendEmail(email, OTP)
	return OTP
}

func sendEmail(email, OTP string) {
	/* Sender data */
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")

	to := []string{email}

	/* smtp server configuration */
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(OTP+" is your otp"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getRandomNum() (string, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(8999))
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(num.Int64()+1000, 10), nil
}

func OTPValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var USEROTP struct {
			Email string
			OTP   string
		}
		if err := c.ShouldBindJSON(&USEROTP); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		var user models.User
		if err := db.DB.Where("email = ? and otp = ?", USEROTP.Email, USEROTP.OTP).First(&user); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong OTP Entered"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "New User Successfully Registered"})
	}
}

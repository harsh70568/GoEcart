package models

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique"`
	FirstName   string    `json:"first_name" gorm:"not null" validate:"required,min=2,max=50"`
	LastName    string    `json:"last_name" gorm:"not null" validate:"required,min=2,max=50"`
	Email       string    `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password    string    `json:"password" gorm:"not null" validate:"required"`
	PhoneNumber string    `json:"phone_no" gorm:"not null;unique" validate:"required"`
	IsAdmin     bool      `json:"isAdmin" gorm:"default:false"`
	OTP         string    `json:"otp"`
	IsBlocked   bool      `json:"isBlocked" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Addresses []Address `gorm:"foreignKey:UserID"` // one to many relationship
}

type Address struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint   `gorm:"not null"`
	Name       string `json:"name" gorm:"not null"`
	Phoneno    string `json:"phoneno" gorm:"not null"`
	Houseno    string `json:"houseno" gorm:"not null"`
	Area       string `json:"area" gorm:"not null"`
	Landmark   string `json:"landmark" gorm:"not null"`
	City       string `json:"city" gorm:"not null"`
	Pincode    string `json:"pincode" gorm:"not null"`
	District   string `json:"district" gorm:"not null"`
	State      string `json:"state" gorm:"not null"`
	Country    string `json:"country" gorm:"not null"`
	Defaultadd bool   `json:"defaultadd" gorm:"default:false"`
}

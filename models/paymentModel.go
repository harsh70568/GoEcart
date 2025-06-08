package models

import "time"

type Payment struct {
	PaymentId     uint      `json:"payment_id" gorm:"primaryKey;autoIncrement"`
	UserId        uint      `json:"user_id"`
	User          User      `gorm:"foreignKey:UserId"`
	PaymentMethod string    `json:"payment_method" gorm:"not null"`
	TotalAmount   float64   `json:"total_amount" gorm:"not null"`
	Status        string    `json:"status" gorm:"not null"`
	Date          time.Time `json:"date"`
}

type OrderDetails struct {
	OrderDetailsID uint      `json:"order_details_id" gorm:"primaryKey;autoIncrement"`
	UserId         uint      `json:"user_id"`
	User           User      `gorm:"foreignKey:UserId"`
	AddressId      uint      `json:"address_id"`
	Address        Address   `gorm:"foreignKey:AddressId"`
	PaymentId      uint      `json:"payment_id"`
	Payment        Payment   `gorm:"foreignKey:PaymentId"`
	ProductId      uint      `json:"product_id"`
	Product        Product   `gorm:"foreignKey:ProductId"`
	OrderItemID    uint      `json:"order_item_id"`
	Quantity       uint      `json:"quantity" gorm:"not null"`
	Status         string    `json:"status" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Coupon struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CouponCode    string    `json:"coupon_code" gorm:"unique;not null"`
	DiscountPrice float64   `json:"discount_price" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at"`
	Expired       time.Time `json:"expired"`
}

type RazorPay struct {
	RazorPaymentID  uint    `json:"razor_payment_id" gorm:"primaryKey;autoIncrement"`
	UserID          uint    `json:"user_id"`
	User            User    `gorm:"foreignKey:UserID"`
	RazorPayOrderID string  `json:"razorpay_order_id" gorm:"not null"`
	Signature       string  `json:"signature" gorm:"not null"`
	AmountPaid      float64 `json:"amount_paid" gorm:"not null"`
}

type Wallet struct {
	ID     uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint    `json:"user_id"`
	User   User    `gorm:"foreignKey:UserID"`
	Amount float64 `json:"amount" gorm:"not null"`
}

type WalletHistory struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID          uint      `json:"user_id"`
	User            User      `gorm:"foreignKey:UserID"`
	Amount          float64   `json:"amount" gorm:"not null"`
	TransactionType string    `json:"transaction_type" gorm:"not null"`
	Date            time.Time `json:"date"`
}

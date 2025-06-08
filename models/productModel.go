package models

type Category struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	CategoryName string `json:"category_name" gorm:"not null"`
}

type Brand struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	BrandName string `json:"brand_name" gorm:"not null"`
}

type Product struct {
	ProductID   uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductName string   `json:"product_name" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
	Stock       uint     `json:"stock" gorm:"not null"`
	Price       uint     `json:"price" gorm:"not null"`
	CategoryId  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryId"`
	BrandId     uint     `json:"brand_id"`
	Brand       Brand    `gorm:"foreignKey:BrandId"`
}

type Cart struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId     uint    `json:"user_id"`
	User       User    `gorm:"foreignKey:UserId"`
	ProductId  uint    `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductId"`
	Quantity   uint    `json:"quantity" gorm:"not null"`
	Price      uint    `json:"price" gorm:"not null"`
	TotalPrice uint    `json:"total_price" gorm:"not null"`
}

type Image struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductId uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId"`
	Image     string  `json:"image" gorm:"not null"`
}

type Wishlist struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    uint    `json:"user_id"`
	User      User    `gorm:"foreignKey:UserId"`
	ProductId uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId"`
}

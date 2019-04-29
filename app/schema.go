package app

type Category struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Enable bool `gorm:"not null" json:"enable"`

}

type Product struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Enable bool `gorm:"not null" json:"enable"`
}

type Image struct {
	ID uint `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
	File string `gorm:"not null" json:"file"`
	Enable bool `gorm:"not null" json:"enable"`
}


type ProductImage struct {
	ProductID uint `gorm:"not null" json:"product_id"`
	ImageID   uint `gorm:"not null" json:"image_id"`
}

type CategoryProduct struct {
	ProductID  uint `gorm:"not null" json:"product_id"`
	CategoryID uint `gorm:"not null" json:"category_id"`

}

type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type CategoryProductResponse struct {
	Product
	Category []Category `json:"categories"`
}

type ProductImagesResponse struct {
	Product
	Image []Image `json:"images"`
}
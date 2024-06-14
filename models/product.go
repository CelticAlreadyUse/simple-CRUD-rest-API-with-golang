package models

type Product struct {
	Id int64  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(255)" json:"product_name"`
	ProductDescription string `gorm:"type:text" json:"product_description"`


}
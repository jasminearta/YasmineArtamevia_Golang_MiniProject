// models/product_log.go
package models

type ProductLog struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	UserID      int    `gorm:"not null" json:"userid"`
	ProductName string `gorm:"not null" json:"productname"`
	Material    string `json:"material"`
	IsPlastic   bool   `json:"isplastic"`
	Rekomendasi string `json:"rekomendasi"` // New field for AI description
}

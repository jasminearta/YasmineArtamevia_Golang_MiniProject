package models

type ProductLog struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	UserID      int    `gorm:"not null" json:"userid"`
	ProductName string `gorm:"not null" json:"productname"`
	Material    string `json:"material"`
	IsPlastic   bool   `json:"isplastic"`
}

package models

type ProductLog struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `gorm:"not null" json:"userid"`
	ProductName string `gorm:"not null" json:"productname"`
	Material    string `json:"material"`
	IsPlastic   bool   `json:"isplastic"`
}

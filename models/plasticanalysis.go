package models

import "time"

type PlasticAnalysis struct {
	ID                   uint `gorm:"primaryKey"`
	ProductID            uint `gorm:"not null"`
	PlasticPercentage    float64
	NonPlasticPercentage float64
	Recommendations      string
	TimeAnalysis         time.Time
}

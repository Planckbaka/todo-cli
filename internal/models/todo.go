package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string   `gorm:"type:text;not null"`
	Description string   `gorm:"type:text"`
	DueDate     string   `gorm:"type:text"`
	Completed   bool     `gorm:"type:boolean;default:false"`
	Priority    string   `gorm:"type:text;default:low"`
	Tags        []string `gorm:"serializer:json;type:text"`
}

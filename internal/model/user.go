package model

type User struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Mobile string `gorm:"uniqueIndex" json:"mobile"`
}
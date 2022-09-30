package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validate:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"  gorm:"type:varchar(100)"`
	Description string `json:"description" binding:"max=20"  gorm:"type:varchar(200)"`
	URL         string `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person `json:"author" binding:"required"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

package models

import "gorm.io/gorm"

type User struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&User{})
}

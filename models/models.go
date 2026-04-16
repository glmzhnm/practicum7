package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type FavoriteBook struct {
	UserID    uint      `gorm:"primaryKey"`
	BookID    uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

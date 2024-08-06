package user

import "time"

type User struct {
	ID           uint      `gorm:"primary_key;autoIncrement"`
	Username     string    `gorm:"column:username"`
	Email        string    `gorm:"column:email;unique_index"`
	PasswordHash string    `gorm:"column:password;not null"`
	Role         string    `gorm:"column:role"`
	IsActive     bool      `gorm:"column:is_active"`
	AddedDate    string    `gorm:"column:added_date"`
	UpdatedDate  string    `gorm:"column:updated_date"`
	Salt         string    `gorm:"column:salt"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	Token        string    `gorm:"column:token"`
}

package model

type User struct {
	UserId uint `gorm:"primaryKey" json:"user_id"`
}

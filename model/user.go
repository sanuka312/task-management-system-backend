package model

type User struct {
	UserId       uint   `gorm:"primaryKey" json:"user_id"`
	UserName     string `'gorm:"size:200;not null" json:"user_name"`
	UserEmail    string `gorm:"size:250;not null" json:"user_email"`
	UserPassword string `gorm:"size:10;not null" json:"user_password"`

	Projects []Project `gorm:"foreignKey:UserId"`
}

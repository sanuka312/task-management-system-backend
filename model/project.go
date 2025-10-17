package model

type Project struct {
	ProjectId          uint   `gorm:"primaryKey" json:"project_id"`
	ProjectName        string `gorm:"size:200;not null" json:"project_name"`
	ProjectDescription string `gorm:"size:200;not null" json:"project_description"`
	UserId             uint   `json:"user_id"` //Foreign key

	User User `gorm:"foreignKey:UserId"`
}

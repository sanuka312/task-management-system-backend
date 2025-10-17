package model

import "time"

type Task struct {
	TaskId          uint      `gorm:"primaryKey" json:"task_id"`
	TaskName        string    `gorm:"size:200;not null" json:"task_name"`
	TaskDescription string    `gorm:"size:400;not null" json:"task_description"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	TaskStatus      string    `gorm:"size:50;default:'pending'" json:"task_status"`
	TaskDueDate     string    `json:"due_date"`

	CompletedAt *time.Time `json:"completed_at"`

	//Foreign keys
	ProjectId uint `json:"project_id"`
}

package model

import "time"

type Task struct {
	ID			unit 		`json:"id" gorm:"primaryKey"`
	Titles		string 		`json:"title" gorm:"not null"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	User		User 		`json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId		unit		`json:"user_id" gorm:"not null"`
}
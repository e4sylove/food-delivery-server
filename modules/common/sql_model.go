package common

import "time"

type SQLModel struct {
	Id int `json:"-" gorm:"column:id;"`
	FakeId *UID `json:"id" gorm:"-"`
	Status int `json:"status" gorm:"status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at;"`
}

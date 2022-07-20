package usermodel

type UserCreate struct {
	UserName string `json:"user_name" gorm:"user_name;"`
	
}
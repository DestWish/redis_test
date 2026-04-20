package models

type User struct {
	ID    int  `json:"id" gorm:"ID" redis:"id"` 
	Name  string `json:"name" gorm:"name" redis:"name"`
	Email string `json:"email" gorm:"email" redis:"email"`
}

type Create_userRequest struct {
	Name string
	Email string
}
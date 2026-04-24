package models

type User struct {
	ID    uint  `json:"id" gorm:"ID" redis:"id"` 
	Name  string `json:"name" gorm:"name" redis:"name"`
	Email string `json:"email" gorm:"email" redis:"email"`
}

type CreateUserRequest struct {
	Name string
	Email string
}
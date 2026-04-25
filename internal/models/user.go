package models

type User struct {
	ID    uint  `json:"id" gorm:"ID" redis:"id"` 
	Name  string `json:"name" gorm:"name" redis:"name"`
	Email string `json:"email" gorm:"email" redis:"email"`
}


type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type ReadUserRequest struct {
	ID uint `json:"id"`
}
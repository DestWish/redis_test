package models

type User struct {
	Login    string   `json:"login" gorm:"login primaryKey" bson:"_id" redis:"login"`
	Name  string `json:"name" gorm:"name" bson:"name" redis:"name"`
	Email string `json:"email" gorm:"email" bson:"email" redis:"email"`
}

type CreateUserRequest struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ReadUserRequest struct {
	Login string `json:"login"`
}

type UpdateUserRequest struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PatchUserRequest struct {
	Login string `json:"login"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type DeleteUserRequest struct {
	Login string `json:"login"`
}

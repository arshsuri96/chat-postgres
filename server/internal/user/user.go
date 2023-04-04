package user

import "context"

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	email    string `json:"email" db:"email"`
	Password string `json:"password" db:"email"`
}

type CreateUserReq struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	email    string `json:"email" db:"email"`
	Password string `json:"password" db:"email"`
}

type CreateUserRes struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	email    string `json:"email" db:"email"`
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
}

type Repository interface {
	createUser(ctx context.Context, user *User) (*User, error)
}

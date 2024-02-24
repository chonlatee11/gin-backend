package domain

import (
	"database/sql"
	"time"

	db "github.com/chonlatee11/gin-backend/db/sqlc"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID    `json:"id"`
	Email    string       `json:"email"`
	CreateAt time.Time    `json:"create_at"`
	UpdateAt sql.NullTime `json:"update_at"`
}


func NewUserResponse(user db.CreateUserRow) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}



type LoginUserResponse struct {
	User  UserResponse `json:"user"`
	// AccessToken string `json:"token"`
}

func NewLoginUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}
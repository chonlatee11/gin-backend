package api

import (
	"net/http"

	"github.com/chonlatee11/gin-backend/api/domain"
	db "github.com/chonlatee11/gin-backend/db/sqlc"
	"github.com/chonlatee11/gin-backend/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) createUser(ctx *gin.Context) {
	var req domain.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassWord(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := domain.NewUserResponse(user)

	ctx.JSON(http.StatusOK, res)
}

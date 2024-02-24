package api

import (
	"net/http"

	"github.com/chonlatee11/gin-backend/api/domain"
	"github.com/chonlatee11/gin-backend/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) loginUser(ctx *gin.Context) {
	var req domain.LoginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	isMatch := util.CheckPasswordHash(req.Password, user.Password)
	if !isMatch {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	// token, err := util.GenerateToken(user.ID, user.Email)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	res := domain.LoginUserResponse{
		User: domain.NewLoginUserResponse(user),
		// AccessToken: token,
	}

	ctx.JSON(http.StatusOK, res)
}

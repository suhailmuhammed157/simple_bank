package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"password"`
}

func (server *Server) Login(ctx *gin.Context) {

	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	validateRes := utils.ValidatePassword(user.HashedPassword, req.Password)

	if !validateRes {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("username or password not found")))
		return
	}

	userResponse := MakeUserResponse(user)

	response := &LoginResponse{
		AccessToken: "",
		User:        userResponse,
	}
	ctx.JSON(http.StatusOK, response)

}

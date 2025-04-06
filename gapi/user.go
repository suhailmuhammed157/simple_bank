package gapi

import (
	"context"

	"github.com/suhailmuhammed157/simple_bank/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (server *Server) Login(context.Context, *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

// func (server *Server) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

// 	hashedPassword, err := utils.EncryptPassword(req.GetPassword())
// 	if err != nil {
// 		er := status.Errorf(codes.Internal, )
// 		return nil, er
// 	}

// 	args := db_source.CreateUserParams{
// 		Username:       req.Username,
// 		HashedPassword: hashedPassword,
// 		FullName:       req.Fullname,
// 		Email:          req.Email,
// 	}

// 	user, err := server.store.CreateUser(ctx, args)
// 	if err != nil {
// 		if pqError, ok := err.(*pq.Error); ok {
// 			switch pqError.Code.Name() {

// 			case "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	formattedUserResponse := MakeUserResponse(user)

// 	ctx.JSON(http.StatusOK, formattedUserResponse)

// }

// func (server *Server) Login(ctx *gin.Context) {

// 	// 	var req LoginRequest
// 	// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 	// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 	// 		return
// 	// 	}

// 	// 	user, err := server.store.GetUser(ctx, req.Username)

// 	// 	if err != nil {
// 	// 		if err == sql.ErrNoRows {
// 	// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 	// 			return
// 	// 		}
// 	// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 	// 		return
// 	// 	}

// 	// 	validateRes := utils.ValidatePassword(user.HashedPassword, req.Password)

// 	// 	if !validateRes {
// 	// 		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("username or password not found")))
// 	// 		return
// 	// 	}

// 	// 	access_token, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.TokenDuration)
// 	// 	if err != nil {
// 	// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 	// 		return
// 	// 	}

// 	// 	refresh_token, refreshPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
// 	// 	if err != nil {
// 	// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 	// 		return
// 	// 	}

// 	// 	server.store.CreateSession(ctx, db_source.CreateSessionParams{
// 	// 		ID:           refreshPayload.ID,
// 	// 		Username:     user.Username,
// 	// 		RefreshToken: refresh_token,
// 	// 		UserAgent:    ctx.Request.UserAgent(),
// 	// 		ClientIp:     ctx.ClientIP(),
// 	// 		IsBlocked:    false,
// 	// 		ExpiresAt:    refreshPayload.ExpiredAt,
// 	// 	})

// 	// 	response := &LoginResponse{
// 	// 		AccessToken:           access_token,
// 	// 		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
// 	// 		RefreshToken:          refresh_token,
// 	// 		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
// 	// 		User:                  MakeUserResponse(user),
// 	// 	}
// 	// 	ctx.JSON(http.StatusOK, response)

// }

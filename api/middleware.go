package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suhailmuhammed157/simple_bank/token"
)

const (
	authorizationPayload = "authorization_payload"
)

func AuthenticateUser(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("missing authentication token")))
			ctx.Abort()
			return
		}

		// The token should be prefixed with "Bearer "
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid authentication token")))
			ctx.Abort()
			return
		}

		tokenString = tokenParts[1]

		payload, err := tokenMaker.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			ctx.Abort()
			return
		}

		ctx.Set(authorizationPayload, payload)
		ctx.Next()
	}

}

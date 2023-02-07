package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
)

func GetTokenAndClaims(ctx *gin.Context) (oauth2.TokenInfo, *generates.JWTAccessClaims) {
	token := ctx.MustGet("__API_TOKEN__").(oauth2.TokenInfo)
	claims := ctx.MustGet("__API_CLAIMS__").(*generates.JWTAccessClaims)
	return token, claims
}

package api

import (
	"net/http"

	"PACKAGENAME/core/auth"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	viewerService *auth.ViewerService
}

func NewUserController(viewerService *auth.ViewerService) *UserController {
	return &UserController{
		viewerService: viewerService,
	}
}

func (c *UserController) Viewer(ctx *gin.Context) {
	token, _ := GetTokenAndClaims(ctx)
	uid := token.GetUserID()

	err := c.viewerService.FindViewer(uid)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, c.viewerService.Viewer)
	}

}

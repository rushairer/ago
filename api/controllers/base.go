package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewParams(params gin.H) gin.H {
	if gin.IsDebugging() {
		params["styleVersion"] = time.Now()
	}
	return params
}

func NewJsonError(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

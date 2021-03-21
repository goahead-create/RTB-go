package handler

import (
	"RTB-go/web/app/controller"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.Negotiate(controller.NewIndexController(ctx).Index())
}

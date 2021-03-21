package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type IndexController interface {
	Index() (int, gin.Negotiate)
}

func NewIndexController(ctx *gin.Context) IndexController {
	return &indexController{
		Ctx: ctx,
	}
}

type indexController struct {
	Ctx *gin.Context
}

func (ctr *indexController) Index() (int, gin.Negotiate) {
	response := "ok"

	return 200, gin.Negotiate{
		Offered: []string{binding.MIMEJSON},
		Data:    response,
	}
}

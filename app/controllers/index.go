package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/panupongdeve/gin-web-start/app/controllers/helloController"
)

func Start(router *gin.Engine) {
	helloController.Start(router)
}

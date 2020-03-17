package helloController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panupongdeve/gin-web-start/repositories/helloRepository"
)

func Start(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, helloRepository.GetHello())
	})
}

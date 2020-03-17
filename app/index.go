package app

import (
	"github.com/gin-gonic/gin"
	"github.com/panupongdeve/gin-web-start/app/controllers"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return r
}

func Start() {
	router := setupRouter()
	controllers.Start(router)
	router.Run(":8080")
}

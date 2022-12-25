package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "wayBack/docs"
	h "wayBack/h"
)

// @title Way API
// @version 0.0.1
// @description Whiskeyyyyyydddyy
// @host localhost:8000
// @BasePath /
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/whiskey", h.GetWhiskeysHandler)
	r.Run(":8000")
}

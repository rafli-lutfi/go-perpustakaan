package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/perpustakaan/config"
	"github.com/rafli-lutfi/perpustakaan/utils"
)

func init() {
	config.LoadEnv()
	utils.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Nahlo",
		})
	})

	r.Run()
}

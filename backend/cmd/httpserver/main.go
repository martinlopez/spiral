package main

import (
	"martinlopez/spiral/internal/core/service/spiralsrv"
	"martinlopez/spiral/internal/handler/spiralhdl"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
			AllowHeaders:     []string{"Authorization"},
		},
	))
	spiralSvc := spiralsrv.New()
	spiralHdl := spiralhdl.NewHTTPHandler(spiralSvc)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/spiral", spiralHdl.Get)

	r.Run()
}

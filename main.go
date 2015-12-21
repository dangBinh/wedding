package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wedding/controllers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default()
	hc := controllers.NewHomeControlelr()
	// Static file
	r.Static("/public", "./public")
	r.LoadHTMLGlob("views/*")
	r.RouterGroup.GET("/", hc.Index)
	r.RouterGroup.POST("/bless", hc.Bless)
	r.Run(":" + port)
}

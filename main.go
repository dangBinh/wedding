package main

import (
	"github.com/wedding/Godeps/_workspace/src/github.com/gin-gonic/gin"
	_ "github.com/wedding/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
	"github.com/wedding/controllers"
)

func main() {
	r := gin.Default()
	hc := controllers.NewHomeControlelr()
	// Static file
	r.Static("/public", "./public")
	r.LoadHTMLGlob("views/*")
	r.GET("/", hc.Index)
	r.POST("/bless", hc.Bless)
	r.Run(":8000")
}

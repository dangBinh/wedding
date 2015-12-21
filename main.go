package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wedding/controllers"
)

func main() {
	r := gin.Default()
	hc := controllers.NewHomeControlelr()
	// Static file
	r.Static("/public", "./public")
	r.LoadHTMLGlob("views/*")
	r.RouterGroup.GET("/", hc.Index)
	r.RouterGroup.POST("/bless", hc.Bless)
	r.Run(":8000")
}

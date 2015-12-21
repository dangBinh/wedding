package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/wedding/models"
	"gopkg.in/gorp.v1"
)

type HomeController struct {
}

// HomeController Constructor
func NewHomeControlelr() *HomeController {
	return &HomeController{}
}

// index
func (hc *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Wedding",
	})
}

func (hc *HomeController) Bless(c *gin.Context) {
	var dbmap = initDB()
	var user models.Participant
	c.Bind(&user)
	name := c.PostForm("name")
	email := c.PostForm("email")
	content := c.PostForm("content")
	if name != "" && email != "" && content != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO Participant (Name, Email, Content) VALUES (?, ?, ?)`, name, email, content); insert != nil {
			_, err := insert.LastInsertId()
			if err == nil {
				c.JSON(200, gin.H{
					"status":   "OK",
					"messages": "Thank you!",
				})
			} else {
				c.JSON(201, gin.H{
					"status":   "Fail",
					"messages": "Có lỗi xảy ra!",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"status":   "Fail",
				"messages": "Có lỗi xảy ra!",
			})
		}
	}
}

func convertDataSource(ds string) (result string) {
	url, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", url.User.String(), url.Host, url.Path)
	return result
}

func initDB() *gorp.DbMap {
	var datasource string
	if os.Getenv("DATABASE_URL") != "" {
		datasource = os.Getenv("DATABASE_URL")
		db, err := sql.Open("postgres", datasource)
		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
		dbmap.AddTableWithName(models.Participant{}, "Participant").SetKeys(true, "Id")
		dbmap.CreateTablesIfNotExists()
		if err != nil {
			log.Fatal(err)
		}
		return dbmap
	} else {
		datasource = "root:admin123@/wedding"
		db, err := sql.Open("mysql", datasource)
		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		dbmap.AddTableWithName(models.Participant{}, "Participant").SetKeys(true, "Id")
		dbmap.CreateTablesIfNotExists()
		if err != nil {
			log.Fatal(err)
		}
		return dbmap
	}

}

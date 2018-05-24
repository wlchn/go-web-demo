package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("mysql", "root:@/go_web_demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/articles")
	{
		v1.GET("/", ArticleIndex)
		v1.POST("/", ArticleCreate)
		v1.GET("/:id", ArticleShow)
		v1.PUT("/:id", ArticleUpdate)
		v1.DELETE("/:id", ArticleDelete)
	}

	router.Run()
}

func ArticleIndex(c *gin.Context) {
	c.String(200, "Article Index")
}

func ArticleCreate(c *gin.Context) {
}

func ArticleShow(c *gin.Context) {
}

func ArticleUpdate(c *gin.Context) {
}

func ArticleDelete(c *gin.Context) {
}
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Article struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {

	db, err = gorm.Open("mysql", "root:@/go_web_demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Article{})

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
	var articles []Article
	if err := db.Find(&articles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, articles)
	}
}

func ArticleCreate(c *gin.Context) {
	var article Article
	c.BindJSON(&article)

	db.Create(&article)
	c.JSON(200, article)
}

func ArticleShow(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article

	if err := db.First(&article, id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, article)
	}
}

func ArticleUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article

	if err := db.First(&article, id); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.BindJSON(&article)
	db.Save(&article)
	c.JSON(200, article)
}

func ArticleDelete(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article
	d := db.Where("id = ?", id).Delete(&article)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

package main

import (
	"qiita/controller"
	"qiita/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// middleware
	engine.Use(middleware.RecordUaAndTime)
	// CRUD book
	bookEngin := engine.Group("/book")
	{
		v1 := bookEngin.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete:id", controller.BookDelete)
		}
	}
	engine.Run(":80")
}

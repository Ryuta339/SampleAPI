package controller

import (
	"log"
	"net/http"
	"qiita/model"
	"qiita/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		log.Print(err.Error())
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	if err := bookService.SetBook(&book); err != nil {
		log.Print(err.Error())
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookList(c *gin.Context) {
	bookService := service.BookService{}
	bookLists := bookService.GetBookList()
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"data":   bookLists,
	})
}

func BookUpdate(c *gin.Context) {
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	if err := bookService.UpdateBook(&book); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// c.String(http.StatusBadRequest, "Bad request")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	bookService := service.BookService{}
	if err := bookService.DeleteBook(int(intId)); err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

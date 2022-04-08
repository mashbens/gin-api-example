package controllers

import (
	"net/http"
	"pustaka-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var BookList = make(map[int]models.Book, 0)

var lastID int = 1

func GetAllBooksHandler(c *gin.Context) {
	var result []models.Book

	for k, _ := range BookList {
		var tempBook models.Book
		tempBook = BookList[k]
		result = append(result, tempBook)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "succsess get all books",
		"data":    result,
	})
}

func PostBookByIDHandler(c *gin.Context) {
	req := new(models.BookInput)

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	book := models.Book{
		ID:    lastID,
		Title: req.Title,
		Price: req.Price,
	}
	BookList[lastID] = book
	result := BookList[lastID]
	lastID++

	c.JSON(http.StatusOK, gin.H{
		"message": "succsess post book",
		"data":    result,
	})
}

func GetbooksByIDHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	_ = err
	book := BookList[idInt]
	c.JSON(http.StatusOK, gin.H{
		"message": "succsess get book by id",
		"data":    book,
	})
}

func UpdateBookByIDHandler(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	_ = err

	req := new(models.BookInput)

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	book := models.Book{
		ID:    idInt,
		Title: req.Title,
		Price: req.Price,
	}
	BookList[idInt] = book
	result := BookList[idInt]

	c.JSON(http.StatusOK, gin.H{
		"message": "succsess update book by id",
		"data":    result,
	})
}

func DelleteBookByIDHandler(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	_ = err

	delete(BookList, idInt)

	c.JSON(http.StatusOK, gin.H{
		"message": "succsess delete book by id",
	})
}

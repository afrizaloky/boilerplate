package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/afrizaloky/boilerplate/model"
	"github.com/afrizaloky/boilerplate/service"
	"github.com/gin-gonic/gin"
)

// BookController : represent the book's controller contract
type BookController interface {
	Name() string
	AddBook(*gin.Context)
	GetAllBook(*gin.Context)
	FindByID(*gin.Context)
}

type bookController struct {
	bookService service.BookService
}

// NewBookController -> returns new book controller
func NewBookController(s service.BookService) BookController {
	return bookController{
		bookService: s,
	}
}
func (b bookController) Name() string {
	return "BookController"
}

func (b bookController) GetAllBook(c *gin.Context) {
	log.Printf("[%s]...get all Books", b.Name())

	books, err := b.bookService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (b bookController) AddBook(c *gin.Context) {
	log.Printf("[%s]...add Book", b.Name())
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := b.bookService.Save(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (b bookController) FindByID(c *gin.Context) {
	log.Printf("[%s]...get book by id", b.Name())
	idStr, found := c.Params.Get("id")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting books"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting books"})
		return
	}

	books, err := b.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

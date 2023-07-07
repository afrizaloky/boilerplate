package route

import (
	"log"

	"github.com/afrizaloky/boilerplate/controller"
	"github.com/afrizaloky/boilerplate/repository"
	"github.com/afrizaloky/boilerplate/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	bookRepository := repository.NewBookRepository(db)

	if err := bookRepository.Migrate(); err != nil {
		log.Fatal("Book migrate err", err)
	}
	bookService := service.NewBookService(bookRepository)

	bookController := controller.NewBookController(bookService)

	books := httpRouter.Group("books")

	books.GET("/", bookController.GetAllBook)
	books.POST("/", bookController.AddBook)

	httpRouter.Run()

}

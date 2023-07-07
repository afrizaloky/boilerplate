package service

import (
	"github.com/afrizaloky/boilerplate/model"
	"github.com/afrizaloky/boilerplate/repository"
	"gorm.io/gorm"
)

// BookService : represent the book's service contract
type BookService interface {
	Save(model.Book) (model.Book, error)
	GetAll() ([]model.Book, error)
	FindByID(id int) (model.Book, error)
	WithTrx(*gorm.DB) bookService
}

type bookService struct {
	bookRepository repository.BookRepository
}

// NewBookService -> returns new book service
func NewBookService(r repository.BookRepository) bookService {
	return bookService{
		bookRepository: r,
	}
}

// WithTrx enables repository with transaction
func (b bookService) WithTrx(trxHandle *gorm.DB) bookService {
	b.bookRepository = b.bookRepository.WithTxn(trxHandle)
	return b
}

func (b bookService) Save(book model.Book) (model.Book, error) {

	return b.bookRepository.Save(book)
}

func (b bookService) GetAll() ([]model.Book, error) {

	return b.bookRepository.GetAll()
}

func (b bookService) FindByID(id int) (model.Book, error) {

	return b.bookRepository.FindByID(id)
}

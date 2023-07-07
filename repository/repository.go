package repository

import (
	"log"

	"github.com/afrizaloky/boilerplate/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	Name() string
	Save(book model.Book) (model.Book, error)
	FindByID(id int) (model.Book, error)
	GetAll() ([]model.Book, error)
	WithTxn(*gorm.DB) bookRepository
	Migrate() error
}

type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) bookRepository {
	return bookRepository{
		DB: db,
	}
}
func (b bookRepository) Name() string {
	return "BookRepository"
}

func (b bookRepository) Migrate() error {
	log.Printf("[%s]...Migrate", b.Name())
	return b.DB.AutoMigrate(&model.Book{})
}

func (b bookRepository) Save(book model.Book) (model.Book, error) {
	log.Printf("[%s]...Save", b.Name())
	err := b.DB.Create(&book).Error
	return book, err

}

func (b bookRepository) FindByID(id int) (model.Book, error) {
	book := model.Book{ID: id}

	log.Printf("[%s]...Get All", b.Name())
	err := b.DB.Find(&book).Error
	return book, err

}

func (b bookRepository) GetAll() (books []model.Book, err error) {
	log.Printf("[%s]...Get All", b.Name())
	err = b.DB.Find(&books).Error
	return books, err

}

func (b bookRepository) WithTxn(trxHandle *gorm.DB) bookRepository {
	if trxHandle == nil {
		log.Printf("T%s Database not found", b.Name())
		return b
	}
	b.DB = trxHandle
	return b
}

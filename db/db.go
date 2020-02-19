package db

import (
	"github.com/likejehu/crudapi/models"
)

// BooksDB is mine books
var BooksDB = map[string]*models.Book{}

type booksDB map[string]*models.Book

// DB is mine representation of database
type DB struct {
	*booksDB
}

package db

import (
	"github.com/likejehu/crudapi/models"
)

type booksDB map[string]*models.Book

// DB is mine representation of database
type DB struct {
	*booksDB
}

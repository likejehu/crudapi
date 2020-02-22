package db

import (
	"github.com/likejehu/crudapi/models"
)

// BooksDB is mine books. и я ней не пользуюсь, игорь, но пусть пока будет
var BooksDB = map[string]*models.Book{}

var booksDB = map[string]*models.Book{}

// DB is mine representation of database
type DB struct {
	b map[string]*models.Book
}

// Delete is delete func
func (d *DB) Delete(key string) {
	delete(d.b, key)
	return
}

// Get is read func
func (d *DB) Get(key string) models.Book {
	v := *d.b[key]
	return v
}

// Put is create func
func (d *DB) Put(key string, book *models.Book) {
	d.b[key] = book
	return
}

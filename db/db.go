package db

import (
	"github.com/likejehu/crudapi/models"
)

// DB is mine representation of database
type DB struct {
	b map[string]*models.Book
}

func newDB() *DB {
	var d DB
	d.b = make(map[string]*models.Book)
	return &d
}

// Library is mine library
var Library = newDB()

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

// Post is create func
func (d *DB) Post(key string, book *models.Book) models.Book {
	d.b[key] = book
	return *d.b[key]
}

// Update is update func
func (d *DB) Update(key string, book *models.Book) (v *models.Book) {
	d.b[key] = book
	v = d.b[key]
	return
}

// GetAll is for returning all the books
func (d *DB) GetAll() map[string]*models.Book {

	return d.b
}

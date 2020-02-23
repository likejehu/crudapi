package db

import (
	"github.com/likejehu/crudapi/models"
)

// DB is mine representation of database
type DB struct {
	b map[string]*models.Book
}

// Library is mine library
var Library = new(DB)

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

// Update is update func
func (d *DB) Update(key string, book *models.Book) (v *models.Book) {
	d.b[key] = book
	v = d.b[key]
	return
}

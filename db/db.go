package db

import (
	"github.com/likejehu/crudapi/models"
)

// DB is mine representation of database
type DB struct {
	B map[string]*models.Book
}

func newDB() *DB {
	var d DB
	d.B = make(map[string]*models.Book)
	return &d
}

// Library is mine library
var Library = newDB()

/*
// Delete is delete func
func (d *DB) Delete(key string) {
	delete(d.B, key)
	return
}

// Get is read func
func (d *DB) Get(key string) *models.Book {
	v := d.B[key]
	return v
}

// Post is create func
func (d *DB) Post(key string, book *models.Book) models.Book {
	d.B[key] = book
	return *d.B[key]
}

// Update is update func
func (d *DB) Update(key string, book *models.Book) (v *models.Book) {
	d.B[key] = book
	v = d.B[key]
	return
}

// GetAll is for returning all the books
func (d *DB) GetAll() map[string]*models.Book {
	books := d.B
	return books
}
*/

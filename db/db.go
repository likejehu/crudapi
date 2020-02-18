package db

import (
	"fmt"

	"github.com/likejehu/crudapi/models"
	"github.com/prologic/bitcask"
)

// BooksDB  is mine books storage representation
var BooksDB = map[string]*models.Book{}

// DB is mine representation of database
var DB, _ = bitcask.Open("/tmp/db")

func main() {
	DB, _ := bitcask.Open("/tmp/db")
	defer DB.Close()
	DB.Put([]byte("Hello"), []byte("World"))
	val, _ := DB.Get([]byte("Hello"))
	fmt.Println(val)
}

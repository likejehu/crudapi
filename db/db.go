package db

import (
	"fmt"

	"github.com/prologic/bitcask"
)

// DB is mine representation of database
var DB, _ = bitcask.Open("/tmp/db")

func main() {
	DB, _ := bitcask.Open("/tmp/db")
	defer DB.Close()
	DB.Put([]byte("Hello"), []byte("World"))
	val, _ := DB.Get([]byte("Hello"))
	fmt.Println(val)
}

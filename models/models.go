package models

import "time"

// Book  is mine book representation
type Book struct {
	Title       string
	Author      string
	Publisher   string
	PublishDate time.Time
	Rating      uint8
	Status      string
}

// Books  is mine books storage representation
var Books = make(map[string]*Book)

// V is for vendetta
var V = 1

func main() {

}

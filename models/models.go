package models

import (
	"time"
)

// Book  is mine book representation
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"date"`
	Rating      uint8     `json:"rating"`
	Status      string    `json:"status"`
}

func main() {

}

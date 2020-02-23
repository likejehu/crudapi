package models

// Book  is mine book representation
type Book struct {
	ID          string `json:"id" form:"id" query:"id" validate:"required"`
	Title       string `json:"title" form:"title" query:"title" validate:"required"`
	Author      string `json:"author" form:"author" query:"author" validate:"required"`
	Publisher   string `json:"publisher" form:"publisher" query:"publisher" validate:"required"`
	PublishDate string `json:"date" form:"date" query:"date" validate:"required"`
	Rating      uint8  `json:"rating" form:"rating" query:"rating" validate:"gte=1,lte=3"`
	Status      string `json:"status" form:"status" query:"status" validate:"oneof=CheckedIn, CheckedOut"`
}

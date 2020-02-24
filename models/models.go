package models

type (
	// Book  is mine book representation
	Book struct {
		ID          string `json:"id" form:"id" query:"id" validate:"required"`
		Title       string `json:"title" form:"title" query:"title" validate:"required"`
		Author      string `json:"author" form:"author" query:"author" validate:"required"`
		Publisher   string `json:"publisher" form:"publisher" query:"publisher" validate:"required"`
		PublishDate string `json:"date" form:"date" query:"date" validate:"required"`
		Rating      uint8  `json:"rating" form:"rating" query:"rating" validate:"required"`
		Status      string `json:"status" form:"status" query:"status" validate:"required"`
	}
)

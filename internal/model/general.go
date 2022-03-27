package model

type ListRequest struct {
	PageNumber int `form:"page_number" json:"page_number" binding:"min=1"`
	PageSize   int `form:"page_size" json:"page_size" binding:"min=5"`
}

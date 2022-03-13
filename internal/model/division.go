package model

type CreateDivisionRequest struct {
	Name string `json:"division_name" binding:"required"`
}

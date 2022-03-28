package model

import (
	"time"

	"github.com/google/uuid"
)

type Division struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"division_name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateDivisionRequest struct {
	Name string `json:"division_name" binding:"required"`
}

type DeleteDivisionRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetDivisionRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

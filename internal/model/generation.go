package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateGenerationRequest struct {
	Year string `json:"year" binding:"required,max=4,numeric"`
}

type UpdateGenerationRequest struct {
	GetOrDeleteGenerationRequest
	CreateGenerationRequest
}

type GetOrDeleteGenerationRequest struct {
	GenerationId string `uri:"id" binding:"required,uuid"`
}

type Generation struct {
	ID        uuid.UUID `json:"id"`
	Year      string    `json:"year"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Active    bool      `json:"active"`
}

type ActiveGeneration struct {
	Year string `json:"year"`
}

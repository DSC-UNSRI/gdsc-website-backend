package model

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID          uuid.UUID `json:"id"`
	FullName    string    `json:"fullname"`
	University  string    `json:"university"`
	Role        uuid.UUID `json:"role_id"`
	Division    Division  `json:"division"`
	PicturePath string    `json:"picture_path"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateMemberRequest struct {
	FullName    string `json:"fullname" binding:"required"`
	University  string `json:"university" binding:"required"`
	RoleID      string `json:"role_id" binding:"required,uuid"`
	DivisionID  string `json:"division_id" binding:"uuid"`
	PicturePath string `json:"picture_path"`
}

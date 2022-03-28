package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateRoleRequest struct {
	Name string `json:"role_name" binding:"required,max=255"`
}

type GetOrDeleteRoleRequest struct {
	RoleId string `uri:"id" binding:"required,uuid"`
}

type UpdateRoleRequest struct {
	GetOrDeleteRoleRequest
	CreateRoleRequest
}

type Role struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}

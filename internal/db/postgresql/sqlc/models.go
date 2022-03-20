// Code generated by sqlc. DO NOT EDIT.

package postgresql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CoreTeam struct {
	DivisionID uuid.UUID
	MemberID   uuid.UUID
}

type Division struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

type Member struct {
	ID          uuid.UUID
	FullName    string
	University  string
	RoleID      uuid.UUID
	DivisionID  uuid.NullUUID
	PicturePath sql.NullString
	CreatedAt   time.Time
	DeletedAt   sql.NullTime
}

type Message struct {
	ID          uuid.UUID
	FullName    string
	PhoneNumber string
	Question    string
	CreatedAt   time.Time
	DeletedAt   sql.NullTime
}

type Role struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

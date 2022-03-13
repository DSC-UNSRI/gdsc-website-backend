// Code generated by sqlc. DO NOT EDIT.

package postgresql

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateDivision(ctx context.Context, name string) (Division, error)
	CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error)
	DeleteDivision(ctx context.Context, divisionID uuid.UUID) error
	DeleteMember(ctx context.Context, memberid uuid.UUID) error
	GetDivision(ctx context.Context, divisionid uuid.UUID) (Division, error)
	GetMember(ctx context.Context, memberid uuid.UUID) (Member, error)
	ListDivisions(ctx context.Context, arg ListDivisionsParams) ([]Division, error)
	ListMembers(ctx context.Context, arg ListMembersParams) ([]Member, error)
	UpdateDivision(ctx context.Context, arg UpdateDivisionParams) (Division, error)
	UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error)
}

var _ Querier = (*Queries)(nil)
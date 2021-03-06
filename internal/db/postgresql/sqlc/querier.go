// Code generated by sqlc. DO NOT EDIT.

package postgresql

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateDivision(ctx context.Context, arg CreateDivisionParams) (Division, error)
	CreateGeneration(ctx context.Context, year string) (CreateGenerationRow, error)
	CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error)
	CreateRole(ctx context.Context, name string) (Role, error)
	CreateSetting(ctx context.Context, arg CreateSettingParams) (Setting, error)
	DeleteDivision(ctx context.Context, divisionID uuid.UUID) (int64, error)
	DeleteGeneration(ctx context.Context, id uuid.UUID) (int64, error)
	DeleteMember(ctx context.Context, memberid uuid.UUID) error
	DeleteRole(ctx context.Context, roleid uuid.UUID) (int64, error)
	GetActiveGeneration(ctx context.Context) (GetActiveGenerationRow, error)
	GetDivision(ctx context.Context, divisionid uuid.UUID) (Division, error)
	GetGeneration(ctx context.Context, id uuid.UUID) (GetGenerationRow, error)
	GetMember(ctx context.Context, memberid uuid.UUID) (Member, error)
	GetRole(ctx context.Context, roleid uuid.UUID) (Role, error)
	GetSetting(ctx context.Context, settingtype SettingType) (Setting, error)
	ListActiveDivisions(ctx context.Context, arg ListActiveDivisionsParams) ([]Division, error)
	ListAllDivisions(ctx context.Context, arg ListAllDivisionsParams) ([]Division, error)
	ListGeneration(ctx context.Context, arg ListGenerationParams) ([]ListGenerationRow, error)
	ListMembers(ctx context.Context, arg ListMembersParams) ([]Member, error)
	ListRole(ctx context.Context, arg ListRoleParams) ([]Role, error)
	SetActiveGeneration(ctx context.Context, generationid uuid.NullUUID) (SetActiveGenerationRow, error)
	UpdateDivision(ctx context.Context, arg UpdateDivisionParams) (Division, error)
	UpdateGeneration(ctx context.Context, arg UpdateGenerationParams) (UpdateGenerationRow, error)
	UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateSetting(ctx context.Context, arg UpdateSettingParams) (Setting, error)
}

var _ Querier = (*Queries)(nil)

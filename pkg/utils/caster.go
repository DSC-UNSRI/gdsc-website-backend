package utils

import (
	"database/sql"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/google/uuid"
)

func CastMemberCreateRequestToDBModel(req model.CreateMemberRequest) postgresql.CreateMemberParams {
	isDivIDNotNull := false
	isPicPathNotNull := false

	if req.DivisionID != "" {
		isDivIDNotNull = true
	}

	if req.PicturePath != "" {
		isDivIDNotNull = true
	}

	return postgresql.CreateMemberParams{
		FullName:   req.FullName,
		University: req.University,
		RoleID:     uuid.MustParse(req.RoleID),
		DivisionID: uuid.NullUUID{
			UUID:  uuid.MustParse(req.DivisionID),
			Valid: isDivIDNotNull,
		},
		PicturePath: sql.NullString{
			String: req.PicturePath,
			Valid:  isPicPathNotNull,
		},
	}
}

func CastToMemberResponse(memberDb postgresql.Member, memberDivision postgresql.Division) model.Member {

	return model.Member{
		ID:         memberDb.ID,
		FullName:   memberDb.FullName,
		University: memberDb.University,
		Role:       memberDb.RoleID,
		Division: model.Division{
			ID:        memberDivision.ID,
			Name:      memberDivision.Name,
			CreatedAt: memberDivision.CreatedAt,
		},
		PicturePath: memberDb.PicturePath.String,
	}
}

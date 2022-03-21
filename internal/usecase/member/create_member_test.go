package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"
	"time"

	mock_db "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/__mock__"
	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/member"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateMemberSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_db.NewMockStore(mockCtrl)
	usecase := usecase.NewMemberUsecase(store)
	mockMemberId := uuid.NewString()
	mockDivisionId := uuid.NewString()
	mockRoleId := uuid.NewString()

	member := postgresql.Member{
		ID:         uuid.MustParse(mockMemberId),
		FullName:   "Tegar",
		University: "Univ",
		RoleID:     uuid.MustParse(mockRoleId),
		DivisionID: uuid.NullUUID{
			UUID:  uuid.MustParse(mockDivisionId),
			Valid: true,
		},
		PicturePath: sql.NullString{
			String: "SomePath",
			Valid:  true,
		},
		CreatedAt: time.Now(),
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	dataModel := model.CreateMemberRequest{
		FullName:    "Tegar",
		University:  "Univ",
		RoleID:      mockRoleId,
		DivisionID:  mockDivisionId,
		PicturePath: "SomePath",
	}
	division := postgresql.Division{
		ID:        uuid.MustParse(mockDivisionId),
		Name:      "backend",
		CreatedAt: time.Now(),
	}

	memberResposne := model.Member{
		ID:          uuid.MustParse(mockMemberId),
		FullName:    "Tegar",
		University:  "Univ",
		Role:        uuid.MustParse(mockRoleId),
		Division:    model.Division(division),
		PicturePath: "SomePath",
	}

	memberParam := utils.CastMemberCreateRequestToDBModel(dataModel)

	ctx := context.Background()
	store.EXPECT().GetDivision(ctx, uuid.MustParse(mockDivisionId)).Return(division, nil)
	store.EXPECT().CreateMember(ctx, memberParam).Return(member, nil)

	response := usecase.CreateMember(dataModel)

	messageSuccess := "Member berhasil dibuat"

	require.NotEmpty(t, response)
	require.Equal(t, response.Message, messageSuccess)
	require.Equal(t, response.Status, http.StatusCreated)
	require.Equal(t, memberResposne, response.Data.(map[string]interface{})["member"])
}

func TestCreateMemberFailed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_db.NewMockStore(mockCtrl)
	usecase := usecase.NewMemberUsecase(store)
	mockMemberId := uuid.NewString()
	mockDivisionId := uuid.NewString()
	mockRoleId := uuid.NewString()

	member := postgresql.Member{
		ID:         uuid.MustParse(mockMemberId),
		FullName:   "Tegar",
		University: "Univ",
		RoleID:     uuid.MustParse(mockRoleId),
		DivisionID: uuid.NullUUID{
			UUID:  uuid.MustParse(mockDivisionId),
			Valid: true,
		},
		PicturePath: sql.NullString{
			String: "SomePath",
			Valid:  true,
		},
		CreatedAt: time.Now(),
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	dataModel := model.CreateMemberRequest{
		FullName:    "Tegar",
		University:  "Univ",
		RoleID:      mockRoleId,
		DivisionID:  mockDivisionId,
		PicturePath: "SomePath",
	}
	division := postgresql.Division{
		ID:        uuid.MustParse(mockDivisionId),
		Name:      "backend",
		CreatedAt: time.Now(),
	}

	// memberResposne := model.Member{
	// 	ID:          uuid.MustParse(mockMemberId),
	// 	FullName:    "Tegar",
	// 	University:  "Univ",
	// 	Role:        uuid.MustParse(mockRoleId),
	// 	Division:    model.Division(division),
	// 	PicturePath: "SomePath",
	// }

	memberParam := utils.CastMemberCreateRequestToDBModel(dataModel)

	ctx := context.Background()
	store.EXPECT().GetDivision(ctx, uuid.MustParse(mockDivisionId)).Return(division, nil)
	store.EXPECT().CreateMember(ctx, memberParam).Return(member, errors.New("Some Errors"))

	response := usecase.CreateMember(dataModel)

	messageSuccess := "Gagal membuat member"

	require.NotEmpty(t, response)
	require.Equal(t, response.Message, messageSuccess)
	require.Equal(t, response.Status, http.StatusInternalServerError)
	require.Equal(t, nil, response.Data)
}

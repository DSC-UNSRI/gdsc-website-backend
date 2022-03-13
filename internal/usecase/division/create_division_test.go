package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	mock_db "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/__mock__"
	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateDivisionSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_db.NewMockStore(mockCtrl)
	usecase := usecase.NewDivisionUsecase(store)

	now := time.Now()
	division := postgresql.Division{
		ID:        uuid.New(),
		Name:      "test",
		CreatedAt: now,
	}

	ctx := context.Background()
	store.EXPECT().CreateDivision(ctx, "test").Return(division, nil)

	response := usecase.CreateDivision(model.CreateDivisionRequest{
		Name: division.Name,
	})

	messageSuccess := "Divisi berhasil dibuat"

	require.NotEmpty(t, response)
	require.Equal(t, response.Message, messageSuccess)
	require.Equal(t, response.Status, http.StatusCreated)
	require.Equal(t, division, response.Data.(map[string]interface{})["division"])
}

func TestCreateDivisionLongName(t *testing.T) {
	name := utils.RandomString(300)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_db.NewMockStore(mockCtrl)
	store.EXPECT().CreateDivision(context.Background(), name).Return(postgresql.Division{
		Name: name[:255],
	}, nil)

	usecase := usecase.NewDivisionUsecase(store)
	res := usecase.CreateDivision(model.CreateDivisionRequest{Name: name})
	divisionRes := res.Data.(map[string]interface{})["division"].(postgresql.Division)
	require.Equal(t, name[:255], divisionRes.Name)
	require.Len(t, name[:255], len(divisionRes.Name))
}

func TestCreateDivisionFailed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_db.NewMockStore(mockCtrl)
	var division postgresql.Division
	store.EXPECT().CreateDivision(context.Background(), "test").Return(division, errors.New("something went wrong"))

	usecase := usecase.NewDivisionUsecase(store)
	response := usecase.CreateDivision(model.CreateDivisionRequest{
		Name: "test",
	})
	message := "Gagal membuat divisi"

	require.NotEmpty(t, response)
	require.Equal(t, message, response.Message)
	require.Equal(t, http.StatusInternalServerError, response.Status)
	require.Empty(t, response.Data)
}

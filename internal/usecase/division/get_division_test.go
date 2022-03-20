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
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetDivisionSuccess(t *testing.T) {

	uuid := uuid.New()
	division := postgresql.Division{
		ID:        uuid,
		Name:      "Backend Test",
		CreatedAt: time.Now(),
	}

	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	storeMock := mock_db.NewMockStore(gomockCtrl)

	storeMock.EXPECT().GetDivision(context.Background(), uuid).Return(division, nil)

	usecase := usecase.NewDivisionUsecase(storeMock)
	res := usecase.GetDivision(model.GetDivisionRequest{
		ID: uuid.String(),
	})
	message := "Sukses mengambil data divisi"
	divisionRes := res.Data.(map[string]interface{})["division"]

	require.NotEmpty(t, res)
	require.Equal(t, http.StatusOK, res.Status)
	require.Equal(t, message, res.Message)
	require.Equal(t, division, divisionRes)

}

func TestGetDivisionFailed(t *testing.T) {
	uuid := uuid.New()

	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	storeMock := mock_db.NewMockStore(gomockCtrl)

	response := model.WebServiceResponse{
		Message: "Gagal mengambil data divisi",
		Status:  http.StatusInternalServerError,
		Data:    nil,
	}

	storeMock.EXPECT().GetDivision(context.Background(), uuid).Return(postgresql.Division{}, errors.New("DB nya error"))

	usecase := usecase.NewDivisionUsecase(storeMock)
	res := usecase.GetDivision(model.GetDivisionRequest{
		ID: uuid.String(),
	})

	require.Equal(t, response, res)
}

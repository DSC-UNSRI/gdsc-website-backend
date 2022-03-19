package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	mock_db "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/__mock__"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestDeleteDivisionSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_db.NewMockStore(mockCtrl)
	usecase := usecase.NewDivisionUsecase(store)
	mockuuid := uuid.NewString()

	ctx := context.Background()

	store.EXPECT().DeleteDivision(ctx, uuid.MustParse(mockuuid)).Return(nil)

	response := usecase.DeleteDivision(model.DeleteDivisionRequest{
		ID: mockuuid,
	})

	messageSuccess := "Divisi berhasil dihapus"

	require.NotEmpty(t, response)
	require.Equal(t, response.Message, messageSuccess)
	require.Equal(t, response.Status, http.StatusOK)
	require.Equal(t, nil, response.Data)
}

func TestDeleteDivisionFailed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	store := mock_db.NewMockStore(mockCtrl)
	usecase := usecase.NewDivisionUsecase(store)
	mockuuid := uuid.NewString()

	ctx := context.Background()

	store.EXPECT().DeleteDivision(ctx, uuid.MustParse(mockuuid)).Return(errors.New("Salah UUID"))

	response := usecase.DeleteDivision(model.DeleteDivisionRequest{
		ID: mockuuid,
	})

	messageSuccess := "Gagal menghapus divisi"

	require.NotEmpty(t, response)
	require.Equal(t, response.Message, messageSuccess)
	require.Equal(t, response.Status, http.StatusInternalServerError)
	require.Equal(t, nil, response.Data)
}

package delivery

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	mock_division "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division/__mock__"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateDivisionSuccess(t *testing.T) {
	division := postgresql.Division{
		ID:        uuid.New(),
		Name:      "backend",
		CreatedAt: time.Now(),
	}
	usecaseReturn := model.WebServiceResponse{
		Message: "Divisi berhasil dibuat",
		Status:  201,
		Data: map[string]interface{}{
			"division": division,
		},
	}
	tempRes, err := json.Marshal(usecaseReturn)
	require.NoError(t, err)
	dataResponse := string(tempRes)
	dataModel := model.CreateDivisionRequest{
		Name: division.Name,
	}
	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	usecase := mock_division.NewMockDivisionUsecase(gomockCtrl)
	usecase.EXPECT().CreateDivision(dataModel).Return(usecaseReturn)
	delivery := NewDivisionDelivery(usecase)
	router.POST("/api/v1/divisions", delivery.CreateDivision)

	bodyReq, err := json.Marshal(map[string]interface{}{
		"division_name": division.Name,
	})
	require.NoError(t, err)
	testRequest, err := http.NewRequest(http.MethodPost, "/api/v1/divisions", bytes.NewBuffer(bodyReq))
	require.NoError(t, err)

	router.ServeHTTP(rr, testRequest)

	require.Equal(t, http.StatusCreated, rr.Code)
	body, err := ioutil.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	require.Equal(t, dataResponse, string(body))
}

package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	mock_division "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division/__mock__"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetDivisionSuccess(t *testing.T) {
	rr, router := getRouter()

	uuid := uuid.New()

	message := "Sukses mengambil data divisi"
	divisionMock := model.Division{
		ID:        uuid,
		Name:      "backend",
		CreatedAt: time.Now(),
	}
	mockRes := model.WebServiceResponse{
		Message: message,
		Status:  http.StatusOK,
		Data: map[string]interface{}{
			"division": divisionMock,
		},
	}

	req, err := http.NewRequest(http.MethodGet, "/api/v1/divisions/"+uuid.String(), nil)

	gomockCtrl := gomock.NewController(t)
	usecaseMock := mock_division.NewMockDivisionUsecase(gomockCtrl)
	usecaseMock.EXPECT().GetDivision(model.GetDivisionRequest{
		ID: uuid.String(),
	}).Return(mockRes)

	delivery := NewDivisionDelivery(usecaseMock)
	router.GET("/api/v1/divisions/:id", delivery.GetDivision)

	router.ServeHTTP(rr, req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rr.Code)

	body, err := ioutil.ReadAll(rr.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rr.Code)

	jsonResponse, err := json.Marshal(mockRes)
	require.NoError(t, err)
	bodyString := string(body)
	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, string(jsonResponse), bodyString)
}

func TestGetDivisionValidationError(t *testing.T) {
	rr, router := getRouter()

	uuid := "asdausd"

	gomockCtrl := gomock.NewController(t)
	usecase := mock_division.NewMockDivisionUsecase(gomockCtrl)
	delivery := NewDivisionDelivery(usecase)

	httpReq, err := http.NewRequest(http.MethodGet, "/api/v1/divisions/"+uuid, nil)
	require.NoError(t, err)

	router.GET("/api/v1/divisions/:id", delivery.GetDivision)

	router.ServeHTTP(rr, httpReq)
	require.Equal(t, rr.Code, http.StatusUnprocessableEntity)

	body, err := ioutil.ReadAll(rr.Body)
	require.NoError(t, err)

	res := `{"message":"ID invalid","status":422,"data":null,"errors":["ID invalid"]}`

	require.Equal(t, res, string(body))
}

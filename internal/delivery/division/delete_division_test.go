package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	mock_division "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division/__mock__"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCoba(t *testing.T) {
	rr, router := getRouter()

	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()

	usecase := mock_division.NewMockDivisionUsecase(gomockCtrl)

	// usecase.EXPECT().CreateDivision(model.CreateDivisionRequest{
	// 	Name: "backend",
	// })
	a := "/api/v1/divisions/"
	delivery := NewDivisionDelivery(usecase)
	router.POST(a, delivery.DeleteDivision)

	// bodyReq, err := json.Marshal(map[string]interface{}{
	// 	"id": id,
	// })
	// require.NoError(t, err)
	testRequest, err := http.NewRequest(http.MethodPost, "/api/v1/divisions/", nil)
	require.NoError(t, err)

	router.ServeHTTP(rr, testRequest)

	require.Equal(t, http.StatusOK, rr.Code)
	// body, err := ioutil.ReadAll(rr.Result().Body)
	// require.NoError(t, err)
	// require.Equal(t, dataResponse, string(body))
}
func TestDeleteDivisionSuccess(t *testing.T) {
	rr, router := getRouter()

	usecaseReturn := model.WebServiceResponse{
		Message: "Divisi berhasil dihapus",
		Status:  200,
		Data:    nil,
	}
	tempRes, err := json.Marshal(usecaseReturn)
	require.NoError(t, err)

	dataResponse := string(tempRes)
	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	usecase := mock_division.NewMockDivisionUsecase(gomockCtrl)

	// usecase.EXPECT().CreateDivision(model.CreateDivisionRequest{
	// 	Name: "backend",
	// })

	usecase.EXPECT().DeleteDivision(model.DeleteDivisionRequest{
		ID: "123e4567-e89b-12d3-a456-426655440000",
	}).Return(usecaseReturn)

	URL := "/api/v1/divisions/:id"

	delivery := NewDivisionDelivery(usecase)
	router.DELETE(URL, delivery.DeleteDivision)
	URL = "/api/v1/divisions/123e4567-e89b-12d3-a456-426655440000"
	require.NoError(t, err)
	testRequest, err := http.NewRequest(http.MethodDelete, URL, nil)
	require.NoError(t, err)
	router.ServeHTTP(rr, testRequest)

	require.Equal(t, http.StatusOK, rr.Code)
	body, err := ioutil.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	require.Equal(t, dataResponse, string(body))
}

func TestDeleteDivisionFailedValidation(t *testing.T) {
	rr, router := getRouter()

	usecaseReturn := model.WebServiceResponse{
		Message: "Gagal menghapus divisi",
		Status:  500,
		Data:    nil,
	}
	tempRes, err := json.Marshal(usecaseReturn)
	require.NoError(t, err)

	dataResponse := string(tempRes)
	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	usecase := mock_division.NewMockDivisionUsecase(gomockCtrl)

	// usecase.EXPECT().CreateDivision(model.CreateDivisionRequest{
	// 	Name: "backend",
	// })

	usecase.EXPECT().DeleteDivision(model.DeleteDivisionRequest{
		ID: "123e4567-e89b-12d3-a456-426655440001",
	}).Return(usecaseReturn)

	URL := "/api/v1/divisions/:id"

	delivery := NewDivisionDelivery(usecase)
	router.DELETE(URL, delivery.DeleteDivision)
	URL = "/api/v1/divisions/123e4567-e89b-12d3-a456-426655440001"
	require.NoError(t, err)
	testRequest, err := http.NewRequest(http.MethodDelete, URL, nil)
	require.NoError(t, err)
	router.ServeHTTP(rr, testRequest)

	require.Equal(t, http.StatusInternalServerError, rr.Code)
	body, err := ioutil.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	require.Equal(t, dataResponse, string(body))
}

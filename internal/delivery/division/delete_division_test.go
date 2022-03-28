package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	mock_division "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division/__mock__"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/validations"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

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

func TestDeleteDivisionFailedOnUUIDValidation(t *testing.T) {

	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		require.FailNow(t, "Validator yg di return gin salah tipe datanya")
	}
	validations.InitValidations(validator)

	jsonResponse := `{"message":"ID harus dengan format UUID","status":422,"data":null,"errors":["ID harus dengan format UUID"]}`

	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	usecase := mock_division.NewMockDivisionUsecase(gomockCtrl)
	delivery := NewDivisionDelivery(usecase)

	URL := "/api/v1/divisions/:id"
	router.DELETE(URL, delivery.DeleteDivision)

	URI := "/api/v1/divisions/Bismillah Game Programmer"
	testRequest, err := http.NewRequest(http.MethodDelete, URI, nil)
	require.NoError(t, err)

	router.ServeHTTP(rr, testRequest)

	require.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	body, err := ioutil.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	require.Equal(t, jsonResponse, string(body))
}

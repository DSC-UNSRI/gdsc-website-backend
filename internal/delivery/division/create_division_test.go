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
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/validations"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

func TestCreateDivisionFailedValidation(t *testing.T) {
	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		require.FailNow(t, "Validator yg di return gin salah tipe datanya")
	}

	validations.InitValidations(validator)
	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()

	jsonRequestBody := []byte(`
	{
		"division_name": ""
	}
	`)
	jsonResponse := `{"message":"Nama Divisi tidak boleh kosong","status":422,"data":null,"errors":["Nama Divisi tidak boleh kosong"]}`

	readerBody := bytes.NewReader(jsonRequestBody)
	httpRequest, err := http.NewRequest(http.MethodPost, "/api/v1/divisions", readerBody)
	require.NoError(t, err)

	usecase := mock_division.NewMockDivisionUsecase(gomockCtrl)

	delivery := NewDivisionDelivery(usecase)
	router.POST("/api/v1/divisions", delivery.CreateDivision)

	router.ServeHTTP(rr, httpRequest)

	require.Equal(t, http.StatusUnprocessableEntity, rr.Code)

	actualResponseByte := rr.Body.Bytes()
	require.Equal(t, jsonResponse, string(actualResponseByte))
}

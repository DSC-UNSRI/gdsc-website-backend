package delivery

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	mock_division "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/member/__mock__"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateMemberSuccess(t *testing.T) {
	rr, router := getRouter()
	mockMemberId := uuid.NewString()
	mockDivisionId := uuid.NewString()
	mockRoleId := uuid.NewString()

	division := model.Division{
		ID:        uuid.MustParse(mockDivisionId),
		Name:      "backend",
		CreatedAt: time.Now(),
	}

	member := model.Member{
		ID:         uuid.MustParse(mockMemberId),
		FullName:   "Tegar",
		University: "Univ",
		Role:       uuid.MustParse(mockRoleId),
		Division:   division,
	}

	usecaseReturn := model.WebServiceResponse{
		Message: "Member berhasil dibuat",
		Status:  201,
		Data: map[string]interface{}{
			"member": member,
		},
	}
	tempRes, err := json.Marshal(usecaseReturn)
	require.NoError(t, err)
	dataResponse := string(tempRes)
	dataModel := model.CreateMemberRequest{
		FullName:   "Tegar",
		University: "Univ",
		RoleID:     mockRoleId,
		DivisionID: mockDivisionId,
	}
	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	usecase := mock_division.NewMockMemberUsecase(gomockCtrl)
	usecase.EXPECT().CreateMember(dataModel).Return(usecaseReturn)
	delivery := NewMemberDelivery(usecase)
	router.POST("/api/v1/members", delivery.CreateMember)

	bodyReq, err := json.Marshal(map[string]interface{}{
		"fullname":    member.FullName,
		"university":  member.University,
		"role_id":     member.Role,
		"division_id": member.Division.ID,
	})

	require.NoError(t, err)
	testRequest, err := http.NewRequest(http.MethodPost, "/api/v1/members", bytes.NewBuffer(bodyReq))
	require.NoError(t, err)

	router.ServeHTTP(rr, testRequest)

	require.Equal(t, http.StatusCreated, rr.Code)
	body, err := ioutil.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	require.Equal(t, dataResponse, string(body))
}

func TestCreateMemberFailedOnRequiredValidation(t *testing.T) {
	rr, router := getRouter()
	mockMemberId := uuid.NewString()
	mockDivisionId := uuid.NewString()
	mockRoleId := uuid.NewString()

	division := model.Division{
		ID:        uuid.MustParse(mockDivisionId),
		Name:      "backend",
		CreatedAt: time.Now(),
	}

	member := model.Member{
		ID:         uuid.MustParse(mockMemberId),
		FullName:   "Tegar",
		University: "Univ",
		Role:       uuid.MustParse(mockRoleId),
		Division:   division,
	}

	bodyReq, err := json.Marshal(map[string]interface{}{
		"fullname":    "",
		"university":  member.University,
		"role_id":     member.Role,
		"division_id": member.Division.ID,
	})
	require.NoError(t, err)

	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	usecase := mock_division.NewMockMemberUsecase(gomockCtrl)
	delivery := NewMemberDelivery(usecase)
	router.POST("/api/v1/members", delivery.CreateMember)

	require.NoError(t, err)
	testRequest, err := http.NewRequest(http.MethodPost, "/api/v1/members", bytes.NewBuffer(bodyReq))
	require.NoError(t, err)

	router.ServeHTTP(rr, testRequest)

	jsonResponse := `{"message":"fullname tidak boleh kosong","status":422,"data":null,"errors":["fullname tidak boleh kosong"]}`
	actualResponseByte := rr.Body.Bytes()
	require.Equal(t, jsonResponse, string(actualResponseByte))
}

func TestCreateMemberFailedOnUUIDValidation(t *testing.T) {
	rr, router := getRouter()
	mockMemberId := uuid.NewString()
	mockDivisionId := uuid.NewString()
	mockRoleId := uuid.NewString()

	division := model.Division{
		ID:        uuid.MustParse(mockDivisionId),
		Name:      "backend",
		CreatedAt: time.Now(),
	}

	member := model.Member{
		ID:         uuid.MustParse(mockMemberId),
		FullName:   "Tegar",
		University: "Univ",
		Role:       uuid.MustParse(mockRoleId),
		Division:   division,
	}

	bodyReq, err := json.Marshal(map[string]interface{}{
		"fullname":    member.FullName,
		"university":  member.University,
		"role_id":     "This is Not UUID",
		"division_id": member.Division.ID,
	})
	require.NoError(t, err)

	gomockCtrl := gomock.NewController(t)
	defer gomockCtrl.Finish()
	usecase := mock_division.NewMockMemberUsecase(gomockCtrl)
	delivery := NewMemberDelivery(usecase)
	router.POST("/api/v1/members", delivery.CreateMember)

	require.NoError(t, err)
	testRequest, err := http.NewRequest(http.MethodPost, "/api/v1/members", bytes.NewBuffer(bodyReq))
	require.NoError(t, err)

	router.ServeHTTP(rr, testRequest)

	jsonResponse := `{"message":"role_id harus dengan format UUID","status":422,"data":null,"errors":["role_id harus dengan format UUID"]}`
	actualResponseByte := rr.Body.Bytes()
	require.Equal(t, jsonResponse, string(actualResponseByte))
}

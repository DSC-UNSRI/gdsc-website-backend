package usecase

import (
	"context"
	"fmt"
	"net/http"

	postgresql "github.com/DSC-UNSRI/gdsc-website-backend/internal/db/postgresql/sqlc"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
	"github.com/DSC-UNSRI/gdsc-website-backend/pkg/utils"
)

func (usecase *memberUsecaseImpl) CreateMember(req model.CreateMemberRequest) model.WebServiceResponse {
	memberParam := utils.CastMemberCreateRequestToDBModel(req)

	memberDb, err := usecase.Store.CreateMember(context.Background(), memberParam)
	fmt.Println(memberDb.ID)
	var memberDivision postgresql.Division
	if memberParam.DivisionID.Valid {
		memberDivision, _ = usecase.GetDivision(context.Background(), memberParam.DivisionID.UUID)
	}

	if err != nil {
		return utils.ToWebServiceResponse("Gagal membuat member", http.StatusInternalServerError, nil)
	}

	return utils.ToWebServiceResponse("Member berhasil dibuat", http.StatusCreated, map[string]interface{}{
		"member": utils.CastToMemberResponse(memberDb, memberDivision),
	})
}

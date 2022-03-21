package usecase

import (
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/db"
	"github.com/DSC-UNSRI/gdsc-website-backend/internal/model"
)

//go:generate mockgen -source=./init.go -destination=./__mock__/member.go -package=mock_member
type MemberUsecase interface {
	CreateMember(model.CreateMemberRequest) model.WebServiceResponse
}

var _ MemberUsecase = &memberUsecaseImpl{}

func NewMemberUsecase(store db.Store) MemberUsecase {
	return &memberUsecaseImpl{
		Store: store,
	}
}

type memberUsecaseImpl struct {
	db.Store
}

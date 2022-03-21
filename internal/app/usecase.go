package app

import (
	dUsecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"
	mUsecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/member"
)

type usecases struct {
	division dUsecase.DivisionUsecase
	member   mUsecase.MemberUsecase
}

func (app *App) initUsecase() {
	var usecases usecases
	usecases.division = dUsecase.NewDivisionUsecase(app.store)
	usecases.member = mUsecase.NewMemberUsecase(app.store)
	app.usecase = usecases
}

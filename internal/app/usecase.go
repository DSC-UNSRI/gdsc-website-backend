package app

import usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"

type usecases struct {
	division usecase.DivisionUsecase
}

func (app *App) initUsecase() {
	var usecases usecases
	usecases.division = usecase.NewDivisionUsecase(app.store)
	app.usecase = usecases
}

package app

import (
	division_usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/division"
	generation_usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/generation"
	role_usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/role"
)

type usecases struct {
	division   division_usecase.DivisionUsecase
	generation generation_usecase.GenerationUsecase
	role       role_usecase.RoleUsecase
}

func (app *App) initUsecase() {
	var usecases usecases
	usecases.division = division_usecase.NewDivisionUsecase(app.store)
	usecases.generation = generation_usecase.NewGenerationUsecase(app.store)
	usecases.role = role_usecase.NewGenerationUsecase(app.store)

	app.usecase = usecases
}

package app

import usecase "github.com/DSC-UNSRI/gdsc-website-backend/internal/usecase/user"

type usecases struct {
	userUsecase usecase.UserUsecase
}

func (app *App) initUsecase() {
	var usecases usecases
	usecases.userUsecase = usecase.NewUserUsecase(app.store)
	app.usecase = usecases
}

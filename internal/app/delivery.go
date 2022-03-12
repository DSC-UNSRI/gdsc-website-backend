package app

import delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/user"

type deliveries struct {
	user delivery.UserDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	usecase := app.usecase.userUsecase
	deliveries.user = delivery.NewUserHandler(usecase)

	app.delivery = deliveries
}

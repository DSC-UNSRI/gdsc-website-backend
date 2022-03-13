package app

import delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/division"

type deliveries struct {
	division delivery.DivisionDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	usecase := app.usecase.division
	deliveries.division = delivery.NewDivisionDelivery(usecase)

	app.delivery = deliveries
}

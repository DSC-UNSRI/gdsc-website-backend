package app

import (
	division_delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/division"
	generation_delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/generation"
	role_delivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/role"
)

type deliveries struct {
	division   division_delivery.DivisionDelivery
	generation generation_delivery.GenerationDelivery
	role       role_delivery.RoleDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	deliveries.division = division_delivery.NewDivisionDelivery(app.usecase.division)
	deliveries.generation = generation_delivery.NewGenerationDelivery(app.usecase.generation)
	deliveries.role = role_delivery.NewRoleDelivery(app.usecase.role)

	app.delivery = deliveries
}

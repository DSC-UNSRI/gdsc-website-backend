package app

import (
	dDelivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/division"
	mDelivery "github.com/DSC-UNSRI/gdsc-website-backend/internal/delivery/member"
)

type deliveries struct {
	division dDelivery.DivisionDelivery
	member   mDelivery.MemberDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
	deliveries.division = dDelivery.NewDivisionDelivery(app.usecase.division)
	deliveries.member = mDelivery.NewMemberDelivery(app.usecase.member)

	app.delivery = deliveries
}

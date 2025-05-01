package factories

import (
	"apps/internal/models"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var IncomeFactory = factory.NewFactory(
	&models.Income{},
).Attr("Amount", func(args factory.Args) (interface{}, error) {
	return randomdata.Number(10000), nil
}).Attr("ClientName", func(args factory.Args) (interface{}, error) {
	return randomdata.RandStringRunes(25), nil
}).Attr("ReceivedAt", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	return date, nil
}).SubFactory("User", UserFactory)

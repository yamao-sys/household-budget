package factories

import (
	"apps/internal/models"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var ExpenseFactory = factory.NewFactory(
	&models.Expense{
		Amount: randomdata.Number(10000),
		Category: models.Category(randomdata.Number(0, 8)),
		Description: randomdata.RandStringRunes(100),
	},
).Attr("PaidAt", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	return date, nil
}).SubFactory("User", UserFactory)

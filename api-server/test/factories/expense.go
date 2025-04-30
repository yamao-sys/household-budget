package factories

import (
	"apps/internal/models"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
)

var ExpenseFactory = factory.NewFactory(
	&models.Expense{},
).Attr("Amount", func(args factory.Args) (interface{}, error) {
	return randomdata.Number(10000), nil
}).Attr("Category", func(args factory.Args) (interface{}, error) {
	return models.Category(randomdata.Number(1, 9)), nil
}).Attr("Description", func(args factory.Args) (interface{}, error) {
	return randomdata.RandStringRunes(100), nil
}).Attr("PaidAt", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	return date, nil
}).SubFactory("User", UserFactory)

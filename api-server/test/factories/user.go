package factories

import (
	"apps/internal/models"
	"log"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"golang.org/x/crypto/bcrypt"
)

var CompanyFactory = factory.NewFactory(
	&models.User{
		Name: randomdata.RandStringRunes(8),
		Email: randomdata.Email(),
	},
).Attr("Password", func(args factory.Args) (interface{}, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to generate hash %v", err)
	}
	return string(hash), nil
})

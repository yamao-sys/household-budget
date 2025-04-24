package validators

import (
	"apps/api"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateSignUpCompany(input *api.PostUsersSignUpJSONRequestBody) error {
	return validation.ValidateStruct(input,
		validation.Field(
			&input.Name,
			validation.Required.Error("企業名は必須入力です。"),
			validation.RuneLength(1, 20).Error("企業名は1 ~ 20文字での入力をお願いします。"),
		),
		validation.Field(
			&input.Email,
			validation.Required.Error("Emailは必須入力です。"),
			is.Email.Error("Emailの形式での入力をお願いします。"),
		),
		validation.Field(
			&input.Password,
			validation.Required.Error("パスワードは必須入力です。"),
			validation.Length(8, 24).Error("パスワードは8 ~ 24文字での入力をお願いします。"),
		),
	)
}

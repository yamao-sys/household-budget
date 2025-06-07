package validators

import (
	api "apps/apis"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateIncome(input *api.PostIncomesJSONRequestBody) error {
	return validation.ValidateStruct(input,
		validation.Field(
			&input.Amount,
			validation.Required.Error("金額は必須入力です。"),
		),
		validation.Field(
			&input.ClientName,
			validation.Required.Error("顧客名は必須入力です。"),
		),
	)
}

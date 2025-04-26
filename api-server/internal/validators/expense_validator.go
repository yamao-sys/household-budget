package validators

import (
	"apps/api"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateExpense(input *api.PostExpensesJSONRequestBody) error {
	return validation.ValidateStruct(input,
		validation.Field(
			&input.PaidAt,
			validation.Required.Error("支出日は必須入力です。"),			
		),
		validation.Field(
			&input.Amount,
			validation.Required.Error("金額は必須入力です。"),
		),
		validation.Field(
			&input.Category,
			validation.Required.Error("カテゴリは必須入力です。"),
			// TODO: カテゴリをENUM管理してバリデーション追加する
		),
		validation.Field(
			&input.Description,
			validation.Required.Error("適用は必須入力です。"),
		),
	)
}

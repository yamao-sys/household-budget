import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./request.tsp";
import "./response.tsp";

import "../common";
import "../util";

using Http;
using OpenAPI;

@tag("expenses")
@route("/expenses")
@useAuth([SecuritySchema])
namespace HouseholdBudgetService.Expense {
  interface Root {
    @operationId("get-expenses")
    @summary("Get Expenses")
    @get
    get(
			@query fromDate?: string,
			@query toDate?: string,
		): SuccessResponse<FetchExpenseListsResponse>;

		@operationId("post-expenses")
		@summary("Post Expense")
		@post
		post(
			@body body: StoreExpenseInput,
		): SuccessResponse<StoreExpenseResponse>
			| InternalServerErrorResponse;
  }

	@route("/totalAmounts")
	interface TotalAmounts {
		@operationId("get-expenses-total-amounts")
		@summary("Get Expense Total Amounts")
		@get
		get(
			@query fromDate: string,
			@query toDate: string,
		): SuccessResponse<TotalAmountListsResponse>;
	}

	@route("/categoryTotalAmounts")
	interface CategoryTotalAmounts {
		@operationId("get-expenses-category-total-amounts")
		@summary("Get Expenses Category TotalAmounts")
		@get
		get(
			@query fromDate: string,
			@query toDate: string,
		): SuccessResponse<CategoryTotalAmountListsResponse>;
	}
}

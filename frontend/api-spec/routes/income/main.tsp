import "@typespec/http";
import "@typespec/openapi";
import "@typespec/openapi3";

import "./request.tsp";
import "./response.tsp";

import "../common";
import "../util";

using Http;
using OpenAPI;

@tag("incomes")
@route("/incomes")
@useAuth([SecuritySchema])
namespace HouseholdBudgetService.Income {
  interface Root {
    @operationId("get-incomes")
    @summary("Get Incomes")
    @get
    get(
			@query fromDate?: string,
			@query toDate?: string,
		): SuccessResponse<FetchIncomeListsResponse>;

		@operationId("post-incomes")
		@summary("Post Income")
		@post
		post(
			@body body: StoreIncomeInput,
		): SuccessResponse<StoreIncomeResponse>
			| InternalServerErrorResponse;
  }

	@route("/totalAmounts")
	interface TotalAmounts {
		@operationId("get-incomes-total-amounts")
		@summary("Get Income Total Amounts")
		@get
		get(
			@query fromDate: string,
			@query toDate: string,
		): SuccessResponse<TotalAmountListsResponse>;
	}

	@route("/clientTotalAmounts")
	interface ClientTotalAmounts {
		@operationId("get-incomes-client-total-amounts")
		@summary("Get Income Client TotalAmounts")
		@get
		get(
			@query fromDate: string,
			@query toDate: string,
		): SuccessResponse<ClientTotalAmountListsResponse>;
	}
}

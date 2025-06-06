import "../../models/income.tsp";

alias FetchIncomeListsResponse = IncomeLists;

@summary("Store Income Response")
model StoreIncomeResponse {
	income: Income;
	errors: StoreIncomeValidationError;
}

@summary("Client Total Amount Lists Response")
model ClientTotalAmountListsResponse {
	totalAmounts: ClientTotalAmountLists[];
}

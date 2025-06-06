import "../../models/expense.tsp";

alias FetchExpenseListsResponse = ExpenseLists;

@summary("Total Amount Lists Response")
model TotalAmountListsResponse {
	totalAmounts: TotalAmountLists[];
}

@summary("Category Total Amount Lists Response")
model CategoryTotalAmountListsResponse {
	totalAmounts: CategoryTotalAmountLists[];
}

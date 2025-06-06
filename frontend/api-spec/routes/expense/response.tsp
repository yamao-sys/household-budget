import "../../models/expense.tsp";
import "../../models/util.tsp";

alias FetchExpenseListsResponse = ExpenseLists;

@summary("Category Total Amount Lists Response")
model CategoryTotalAmountListsResponse {
	totalAmounts: CategoryTotalAmountLists[];
}

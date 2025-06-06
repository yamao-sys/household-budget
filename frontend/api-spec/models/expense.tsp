@summary("Expense Lists")
model ExpenseLists {
	expenses: Expense[];
}

@summary("Expense")
model Expense {
	id: string;
	paidAt: plainDate;
	amount: integer;
	category: integer;
	description: string;
}

@summary("Store Expense Response")
model StoreExpenseResponse {
	expense: Expense;
	errors: StoreExpenseValidationError;
}

@summary("Store Expense Validation Error")
model StoreExpenseValidationError{
	paidAt?: string[];
	amount?: string[];
	category?: string[];
	description?: string[];
}

@summary("Total Amount Lists")
model TotalAmountLists {
	date: plainDate;
	extendProps: {
		type: string;
		totalAmount: integer;
	};
}

@summary("Category Total Amount Lists")
model CategoryTotalAmountLists {
	category: integer;
	totalAmount: integer;
}

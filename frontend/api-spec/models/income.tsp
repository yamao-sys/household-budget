@summary("Income")
model Income {
	id: integer;
	receivedAt: plainDate;
	amount: integer;
	clientName: string;
}

@summary("Income Lists")
model IncomeLists {
	incomes: Income[];
}

@summary("Store Income Validation Error")
model StoreIncomeValidationError {
	receivedAt?: string[];
	amount?: string[];
	clientName?: string[];
}

@summary("Client Total Amount Lists")
model ClientTotalAmountLists {
	clientName: string;
	totalAmount: integer;
}

@summary("Store Expense Input")
model StoreExpenseInput {
	paidAt: plainDate;
	amount: integer;
	category: integer;
	description: string;
}

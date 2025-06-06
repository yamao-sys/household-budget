@summary("Total Amount Lists")
model TotalAmountLists {
	date: plainDate;
	extendProps: {
		type: string;
		totalAmount: integer;
	};
}

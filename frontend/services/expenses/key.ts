export const expenseKeys = {
  all: ["expenses"] as const,
  lists: () => [...expenseKeys.all, "list"] as const,
  list: (fromDate: string, toDate: string) => [...expenseKeys.lists(), fromDate, toDate] as const,
  totalAmounts: () => [...expenseKeys.all, "totalAmount"] as const,
  totalAmount: (fromDate: string, toDate: string) => [...expenseKeys.totalAmounts(), fromDate, toDate] as const,
  categoryTotalAmounts: () => [...expenseKeys.all, "categoryTotalAmount"] as const,
  categoryTotalAmount: (fromDate: string, toDate: string) => [...expenseKeys.categoryTotalAmounts(), fromDate, toDate] as const,
} as const;

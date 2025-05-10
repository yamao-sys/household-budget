export const incomeKeys = {
  all: ["incomes"] as const,
  lists: () => [...incomeKeys.all, "list"] as const,
  list: (fromDate: string, toDate: string) => [...incomeKeys.lists(), fromDate, toDate] as const,
  totalAmounts: () => [...incomeKeys.all, "totalAmount"] as const,
  totalAmount: (fromDate: string, toDate: string) => [...incomeKeys.totalAmounts(), fromDate, toDate] as const,
  clientTotalAmounts: () => [...incomeKeys.all, "clientTotalAmount"] as const,
  clientTotalAmount: (fromDate: string, toDate: string) => [...incomeKeys.clientTotalAmounts(), fromDate, toDate] as const,
} as const;

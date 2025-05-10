import { QueryClient, useMutation, useQuery } from "@tanstack/react-query";
import { expenseKeys } from "./key";
import { getExpenses, getExpenseTotalAmounts, postCreateExpense } from "./api";
import type { StoreExpenseInput, StoreExpenseResponse } from "~/types";
import { getDateString } from "~/lib/date";

export const useGetExpenses = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: expenseKeys.list(fromDate, toDate),
    queryFn: () => getExpenses(fromDate, toDate, csrfToken),
  });

  return { data, isPending, isError };
};

export const useGetExpenseTotalAmounts = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: expenseKeys.totalAmount(fromDate, toDate),
    queryFn: () => getExpenseTotalAmounts(fromDate, toDate, csrfToken),
  });

  return { data, isPending, isError };
};

export const usePostCreateExpense = (
  queryClient: QueryClient,
  onMutate: () => void,
  onSuccess: (data: StoreExpenseResponse) => void,
  input: StoreExpenseInput,
  date: string,
  csrfToken: string,
) => {
  return useMutation({
    onMutate: () => onMutate,
    mutationFn: () => postCreateExpense(input, csrfToken),
    onSuccess: (data) => {
      const selectedDate = new Date(date);
      const beginningOfMonth = getDateString(new Date(selectedDate.getFullYear(), selectedDate.getMonth(), 1));
      const endOfMonth = getDateString(new Date(selectedDate.getFullYear(), selectedDate.getMonth() + 1, 0));

      queryClient.invalidateQueries({
        queryKey: expenseKeys.list(date, date),
      });
      queryClient.invalidateQueries({
        queryKey: expenseKeys.totalAmount(beginningOfMonth, endOfMonth),
      });

      onSuccess(data);
    },
  });
};

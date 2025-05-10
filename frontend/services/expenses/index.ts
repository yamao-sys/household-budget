import { QueryClient, useMutation, useQuery } from "@tanstack/react-query";
import { expenseKeys } from "./key";
import { getExpenses, postCreateExpense } from "./api";
import type { StoreExpenseInput, StoreExpenseResponse } from "~/types";

export const useGetExpenses = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: expenseKeys.list(fromDate, toDate),
    queryFn: () => getExpenses(fromDate, toDate, csrfToken),
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
      queryClient.invalidateQueries({
        queryKey: expenseKeys.list(date, date),
      });
      onSuccess(data);
    },
  });
};

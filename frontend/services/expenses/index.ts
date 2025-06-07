import { QueryClient, useMutation, useQuery } from "@tanstack/react-query";
import { getExpenseCategoryTotalAmounts, getExpenses, getExpenseTotalAmounts, postCreateExpense } from "./api";
import { getDateString } from "~/lib/date";
import type { StoreExpenseInput, StoreExpenseResponse } from "~/apis/model";
import { getGetExpensesCategoryTotalAmountsQueryKey, getGetExpensesQueryKey, getGetExpensesTotalAmountsQueryKey } from "~/apis/expenses/expenses";

export const useGetExpenses = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: getGetExpensesQueryKey({ fromDate, toDate }),
    queryFn: () => getExpenses(fromDate, toDate, csrfToken),
  });

  return { data, isPending, isError };
};

export const useGetExpenseTotalAmounts = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: getGetExpensesTotalAmountsQueryKey({ fromDate, toDate }),
    queryFn: () => getExpenseTotalAmounts(fromDate, toDate, csrfToken),
    staleTime: 1000 * 60 * 10, // NOTE: FullCalenderで月を変更すると、キャッシュクリアされてしまうため設定
  });

  return { data, isPending, isError };
};

export const useGetExpenseCategoryTotalAmounts = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: getGetExpensesCategoryTotalAmountsQueryKey({ fromDate, toDate }),
    queryFn: () => getExpenseCategoryTotalAmounts(fromDate, toDate, csrfToken),
    staleTime: 1000 * 60 * 10,
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
        queryKey: getGetExpensesQueryKey({ fromDate: date, toDate: date }),
      });
      queryClient.invalidateQueries({
        queryKey: getGetExpensesTotalAmountsQueryKey({ fromDate: beginningOfMonth, toDate: endOfMonth }),
      });
      queryClient.invalidateQueries({
        queryKey: getGetExpensesCategoryTotalAmountsQueryKey({ fromDate: beginningOfMonth, toDate: endOfMonth }),
      });

      onSuccess(data);
    },
  });
};

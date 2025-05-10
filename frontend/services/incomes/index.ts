import { QueryClient, useMutation, useQuery } from "@tanstack/react-query";
import { incomeKeys } from "./key";
import { getIncomes, getIncomeTotalAmounts, postCreateIncome } from "./api";
import type { StoreIncomeInput, StoreIncomeResponse } from "~/types";
import { getDateString } from "~/lib/date";

export const useGetIncomes = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: incomeKeys.list(fromDate, toDate),
    queryFn: () => getIncomes(fromDate, toDate, csrfToken),
  });

  return { data, isPending, isError };
};

export const useGetIncomeTotalAmounts = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: incomeKeys.totalAmount(fromDate, toDate),
    queryFn: () => getIncomeTotalAmounts(fromDate, toDate, csrfToken),
  });

  return { data, isPending, isError };
};

export const usePostCreateIncome = (
  queryClient: QueryClient,
  onMutate: () => void,
  onSuccess: (data: StoreIncomeResponse) => void,
  input: StoreIncomeInput,
  date: string,
  csrfToken: string,
) => {
  return useMutation({
    onMutate: () => onMutate,
    mutationFn: () => postCreateIncome(input, csrfToken),
    onSuccess: (data) => {
      const selectedDate = new Date(date);
      const beginningOfMonth = getDateString(new Date(selectedDate.getFullYear(), selectedDate.getMonth(), 1));
      const endOfMonth = getDateString(new Date(selectedDate.getFullYear(), selectedDate.getMonth() + 1, 0));

      queryClient.invalidateQueries({
        queryKey: incomeKeys.list(date, date),
      });
      queryClient.invalidateQueries({
        queryKey: incomeKeys.totalAmount(beginningOfMonth, endOfMonth),
      });

      onSuccess(data);
    },
  });
};

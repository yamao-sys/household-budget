import { QueryClient, useMutation, useQuery } from "@tanstack/react-query";
import { incomeKeys } from "./key";
import { getIncomes, getIncomeTotalAmounts, postCreateIncome } from "./api";
import type { StoreIncomeInput, StoreIncomeResponse } from "~/types";

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
      queryClient.invalidateQueries({
        queryKey: incomeKeys.list(date, date),
      });
      onSuccess(data);
    },
  });
};

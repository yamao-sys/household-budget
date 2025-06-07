import { QueryClient, useMutation, useQuery } from "@tanstack/react-query";
import { getIncomeClientTotalAmounts, getIncomes, getIncomeTotalAmounts, postCreateIncome } from "./api";
import { getDateString } from "~/lib/date";
import type { StoreIncomeInput, StoreIncomeResponse } from "~/apis/model";
import { getGetIncomesClientTotalAmountsQueryKey, getGetIncomesQueryKey, getGetIncomesTotalAmountsQueryKey } from "~/apis/incomes/incomes";

export const useGetIncomes = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: getGetIncomesQueryKey({ fromDate, toDate }),
    queryFn: () => getIncomes(fromDate, toDate, csrfToken),
  });

  return { data, isPending, isError };
};

export const useGetIncomeTotalAmounts = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: getGetIncomesTotalAmountsQueryKey({ fromDate, toDate }),
    queryFn: () => getIncomeTotalAmounts(fromDate, toDate, csrfToken),
    staleTime: 1000 * 60 * 10, // NOTE: FullCalenderで月を変更すると、キャッシュクリアされてしまうため設定
  });

  return { data, isPending, isError };
};

export const useGetIncomeClientTotalAmounts = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: getGetIncomesClientTotalAmountsQueryKey({ fromDate, toDate }),
    queryFn: () => getIncomeClientTotalAmounts(fromDate, toDate, csrfToken),
    staleTime: 1000 * 60 * 10,
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
        queryKey: getGetIncomesQueryKey({ fromDate: date, toDate: date }),
      });
      queryClient.invalidateQueries({
        queryKey: getGetIncomesTotalAmountsQueryKey({ fromDate: beginningOfMonth, toDate: endOfMonth }),
      });
      queryClient.invalidateQueries({
        queryKey: getGetIncomesClientTotalAmountsQueryKey({ fromDate: beginningOfMonth, toDate: endOfMonth }),
      });

      onSuccess(data);
    },
  });
};

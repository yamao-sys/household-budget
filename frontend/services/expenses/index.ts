import { useQuery } from "@tanstack/react-query";
import { expenseKeys } from "./key";
import { getExpenses } from "./api";

export const useGetExpenses = (fromDate: string, toDate: string, csrfToken: string) => {
  const { data, isPending, isError } = useQuery({
    queryKey: expenseKeys.list(fromDate, toDate),
    queryFn: () => getExpenses(fromDate, toDate, csrfToken),
  });

  return { data, isPending, isError };
};

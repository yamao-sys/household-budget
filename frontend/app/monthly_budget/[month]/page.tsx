import { getCategoryTotalAmounts } from "~/apis/expenses.api";
import type { Route } from "./+types/page";
import { getDateString } from "~/lib/date";
import { useLoaderData } from "react-router";
import { MonthlyBudgetDetail } from "~/features/monthly-budget/components/Detail/MonthlyBudgetDetail";
import { getClientTotalAmounts } from "~/apis/incomes.api";

export async function clientLoader({ params }: Route.ClientLoaderArgs) {
  const month = params.month;
  if (!/^\d{4}-(0[1-9]|1[0-2])$/.test(month)) {
    throw new Response("Invalid Path", { status: 500 });
  }

  const monthDate = new Date(month);

  const monthBeginningDate = new Date(monthDate.getFullYear(), monthDate.getMonth(), 1);
  const monthEndDate = new Date(monthDate.getFullYear(), monthDate.getMonth() + 1, 0);

  let categoryTotalAmounts = (await getCategoryTotalAmounts(getDateString(monthBeginningDate), getDateString(monthEndDate))) ?? [];
  let clientTotalAmounts = (await getClientTotalAmounts(getDateString(monthBeginningDate), getDateString(monthEndDate))) ?? [];
  return { monthDate, categoryTotalAmounts, clientTotalAmounts };
}

export default function MonthlyBudgetDetailPage() {
  const { monthDate, categoryTotalAmounts, clientTotalAmounts } = useLoaderData<typeof clientLoader>();

  return (
    <>
      <MonthlyBudgetDetail monthDate={monthDate} categoryTotalAmounts={categoryTotalAmounts} clientTotalAmounts={clientTotalAmounts} />
    </>
  );
}

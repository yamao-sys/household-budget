import { getCategoryTotalAmounts } from "~/apis/expenses.api";
import type { Route } from "./+types/page";
import { getDateString } from "~/lib/date";
import { useLoaderData } from "react-router";
import { MonthlyBudgetDetail } from "~/features/monthly-budget/components/Detail/MonthlyBudgetDetail";
import { getClientTotalAmounts } from "~/apis/incomes.api";
import { authContext } from "~/middlewares/auth-context";
import { useAuth } from "~/hooks/useAuth";

export async function clientLoader({ params, context }: Route.ClientLoaderArgs) {
  const auth = context.get(authContext);

  const month = params.month;
  if (!/^\d{4}-(0[1-9]|1[0-2])$/.test(month)) {
    throw new Response("Invalid Path", { status: 500 });
  }

  const monthDate = new Date(month);

  const monthBeginningDate = new Date(monthDate.getFullYear(), monthDate.getMonth(), 1);
  const monthEndDate = new Date(monthDate.getFullYear(), monthDate.getMonth() + 1, 0);

  let categoryTotalAmounts =
    (await getCategoryTotalAmounts(getDateString(monthBeginningDate), getDateString(monthEndDate), auth?.csrfToken ?? "")) ?? [];
  let clientTotalAmounts = (await getClientTotalAmounts(getDateString(monthBeginningDate), getDateString(monthEndDate), auth?.csrfToken ?? "")) ?? [];
  return { isSignedIn: !!auth?.isSignedIn, csrfToken: auth?.csrfToken ?? "", monthDate, categoryTotalAmounts, clientTotalAmounts };
}

export default function MonthlyBudgetDetailPage() {
  const { isSignedIn, csrfToken, monthDate, categoryTotalAmounts, clientTotalAmounts } = useLoaderData<typeof clientLoader>();

  useAuth(isSignedIn, csrfToken);

  return (
    <>
      <MonthlyBudgetDetail monthDate={monthDate} categoryTotalAmounts={categoryTotalAmounts} clientTotalAmounts={clientTotalAmounts} />
    </>
  );
}

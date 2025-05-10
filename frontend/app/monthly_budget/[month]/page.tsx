import type { Route } from "./+types/page";
import { useLoaderData } from "react-router";
import { MonthlyBudgetDetail } from "~/features/monthly-budget/components/Detail/MonthlyBudgetDetail";
import { authContext } from "~/middlewares/auth-context";
import { useAuth } from "~/hooks/useAuth";

export async function clientLoader({ params, context }: Route.ClientLoaderArgs) {
  const auth = context.get(authContext);

  const month = params.month;
  if (!/^\d{4}-(0[1-9]|1[0-2])$/.test(month)) {
    throw new Response("Invalid Path", { status: 500 });
  }

  const monthDate = new Date(month);
  return { isSignedIn: !!auth?.isSignedIn, csrfToken: auth?.csrfToken ?? "", monthDate };
}

export default function MonthlyBudgetDetailPage() {
  const { isSignedIn, csrfToken, monthDate } = useLoaderData<typeof clientLoader>();

  useAuth(isSignedIn, csrfToken);

  return (
    <>
      <MonthlyBudgetDetail monthDate={monthDate} csrfToken={csrfToken} />
    </>
  );
}

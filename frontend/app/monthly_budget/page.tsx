import { MonthlyBudgetCalender } from "~/features/monthly-budget/components/Calender/MonthlyBudgetCalender";
import type { Route } from "./+types/page";
import { authContext } from "~/middlewares/auth-context";
import { useLoaderData } from "react-router";
import { useAuth } from "~/hooks/useAuth";

export async function clientLoader({ context }: Route.ClientLoaderArgs) {
  const auth = context.get(authContext);

  return { isSignedIn: !!auth?.isSignedIn, csrfToken: auth?.csrfToken ?? "" };
}

export default function MonthlyBudgetPage() {
  const { isSignedIn, csrfToken } = useLoaderData<typeof clientLoader>();

  useAuth(isSignedIn, csrfToken);

  return (
    <>
      <MonthlyBudgetCalender />
    </>
  );
}

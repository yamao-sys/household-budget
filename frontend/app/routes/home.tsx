import { authContext } from "~/middlewares/auth-context";
import type { Route } from "./+types/home";
import { useLoaderData, useNavigate } from "react-router";
import { useAuth } from "~/hooks/useAuth";
import { useEffect } from "react";
import { NAVIGATION_PAGE_LIST } from "../routes";

export function meta({}: Route.MetaArgs) {
  return [{ title: "New React Router App" }, { name: "description", content: "Welcome to React Router!" }];
}

export async function clientLoader({ context }: Route.ClientLoaderArgs) {
  const auth = context.get(authContext);

  return { isSignedIn: !!auth?.isSignedIn, csrfToken: auth?.csrfToken ?? "" };
}

export default function Home() {
  const navigate = useNavigate();

  const { isSignedIn, csrfToken } = useLoaderData<typeof clientLoader>();

  useAuth(isSignedIn, csrfToken);

  useEffect(() => {
    const toNavigatePath = isSignedIn ? NAVIGATION_PAGE_LIST.monthlyBudgetPage : NAVIGATION_PAGE_LIST.signInPage;
    navigate(toNavigatePath);
  }, [isSignedIn]);

  return <></>;
}

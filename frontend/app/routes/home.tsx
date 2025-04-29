import { getCheckSignedIn } from "~/apis/users.api";
import type { Route } from "./+types/home";
import { useEffect } from "react";
import { useNavigate } from "react-router";
import { NAVIGATION_PAGE_LIST } from "../routes";

export function meta({}: Route.MetaArgs) {
  return [{ title: "New React Router App" }, { name: "description", content: "Welcome to React Router!" }];
}

export default function Home() {
  const navigate = useNavigate();

  useEffect(() => {
    async function init() {
      const isSignedIn = await getCheckSignedIn();
      const toNavigatePath = isSignedIn ? NAVIGATION_PAGE_LIST.monthlyBudgetPage : NAVIGATION_PAGE_LIST.signInPage;

      navigate(toNavigatePath);
    }
    init();
  }, [navigate]);

  return <></>;
}

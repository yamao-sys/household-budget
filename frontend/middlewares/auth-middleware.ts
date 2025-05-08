import { redirect, type unstable_MiddlewareFunction as MiddlewareFunction } from "react-router";
import { authContext } from "./auth-context";
import { getCsrfToken } from "~/apis/csrf.api";
import { getCheckSignedIn } from "~/apis/users.api";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";

export const authMiddleware: MiddlewareFunction = async ({ request, context }) => {
  // NOTE: 画面遷移の際にCSRFトークンを取得
  await getCsrfToken();

  const checkedSignedIn = await getCheckSignedIn();

  let toNavigatePath = "";
  const url = new URL(request.url);
  const pathname = url.pathname;
  if (pathname === NAVIGATION_PAGE_LIST.signInPage) {
    if (checkedSignedIn) {
      toNavigatePath = NAVIGATION_PAGE_LIST.monthlyBudgetPage;
    }
  }
  if (pathname !== NAVIGATION_PAGE_LIST.signUpPage && pathname !== NAVIGATION_PAGE_LIST.signInPage) {
    if (!checkedSignedIn) {
      toNavigatePath = NAVIGATION_PAGE_LIST.signInPage;
    }
  }
  context.set(authContext, { isSignedIn: checkedSignedIn });

  if (toNavigatePath !== "") {
    throw redirect(toNavigatePath);
  }
};

import { redirect, type unstable_MiddlewareFunction as MiddlewareFunction } from "react-router";
import { authContext } from "./auth-context";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";
import { getCheckSignedIn } from "~/services/users/api";
import { getCsrfToken } from "~/services/csrf/api";

export const authMiddleware: MiddlewareFunction = async ({ request, context }) => {
  // NOTE: 画面遷移の際にCSRFトークンを取得
  const csrfToken = await getCsrfToken();

  const checkedSignedIn = await getCheckSignedIn(csrfToken);

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
  context.set(authContext, { isSignedIn: checkedSignedIn, csrfToken });

  if (toNavigatePath !== "") {
    throw redirect(toNavigatePath);
  }
};

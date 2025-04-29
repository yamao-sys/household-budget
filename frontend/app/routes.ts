import { type RouteConfig, index, layout, route } from "@react-router/dev/routes";

const NAVIGATION_PATH_LIST = {
  top: "/",
  signUpPage: "sign_up",
  signInPage: "sign_in",
  monthlyBudgetPage: "monthly_budget",
};

export const NAVIGATION_PAGE_LIST = {
  top: NAVIGATION_PATH_LIST.top,
  signUpPage: `/${NAVIGATION_PATH_LIST.signUpPage}`,
  signInPage: `/${NAVIGATION_PATH_LIST.signInPage}`,
  monthlyBudgetPage: `/${NAVIGATION_PATH_LIST.monthlyBudgetPage}`,
};

export default [
  layout("./Layout.tsx", [
    index("routes/home.tsx"),
    route(NAVIGATION_PATH_LIST.signUpPage, "sign_up/page.tsx"),
    route(NAVIGATION_PATH_LIST.signInPage, "sign_in/page.tsx"),
    route(NAVIGATION_PATH_LIST.monthlyBudgetPage, "monthly_budget/page.tsx"),
  ]),
] satisfies RouteConfig;

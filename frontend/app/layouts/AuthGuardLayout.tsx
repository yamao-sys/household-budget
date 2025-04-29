import { Outlet, useLocation, useNavigate } from "react-router";
import { getCheckSignedIn } from "~/apis/users.api";
import { NAVIGATION_PAGE_LIST } from "../routes";
import { useEffect } from "react";

export default function AuthGuardLayout() {
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    async function init() {
      const isSignedIn = await getCheckSignedIn();

      if (location.pathname === NAVIGATION_PAGE_LIST.signInPage) {
        if (isSignedIn) {
          navigate(NAVIGATION_PAGE_LIST.monthlyBudgetPage);
          return;
        }
      } else {
        if (!isSignedIn) {
          navigate(NAVIGATION_PAGE_LIST.signInPage);
          return;
        }
      }
    }
    init();
  }, [navigate, location]);

  return (
    <>
      <Outlet />
    </>
  );
}

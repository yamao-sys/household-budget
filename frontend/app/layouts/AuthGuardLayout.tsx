import { Outlet, useLocation, useNavigate } from "react-router";
import { getCheckSignedIn } from "~/apis/users.api";
import { NAVIGATION_PAGE_LIST } from "../routes";
import { useEffect, useState } from "react";
import { HeaderNavigation } from "../HeaderNavigation";

export default function AuthGuardLayout() {
  const navigate = useNavigate();
  const location = useLocation();
  const [isSignedIn, setIsSignedIn] = useState(false);

  useEffect(() => {
    async function init() {
      setIsSignedIn(await getCheckSignedIn());

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
  }, [navigate, location, isSignedIn]);

  return (
    <>
      <HeaderNavigation isSignedIn={isSignedIn}>
        <Outlet />
      </HeaderNavigation>
    </>
  );
}

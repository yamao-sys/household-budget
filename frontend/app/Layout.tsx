import { Outlet, useLocation, useNavigate } from "react-router";
import { getCheckSignedIn } from "~/apis/users.api";
import { NAVIGATION_PAGE_LIST } from "./routes";
import { useEffect, useState } from "react";
import { HeaderNavigation } from "./HeaderNavigation";

export default function Layout() {
  const navigate = useNavigate();
  const location = useLocation();
  const [isSignedIn, setIsSignedIn] = useState(false);

  useEffect(() => {
    async function init() {
      const checkedSignedIn = await getCheckSignedIn();

      let toNavigatePath = "";
      if (location.pathname === NAVIGATION_PAGE_LIST.signInPage) {
        if (checkedSignedIn) {
          toNavigatePath = NAVIGATION_PAGE_LIST.monthlyBudgetPage;
        }
      }
      if (location.pathname !== NAVIGATION_PAGE_LIST.signUpPage && location.pathname !== NAVIGATION_PAGE_LIST.signInPage) {
        if (!checkedSignedIn) {
          toNavigatePath = NAVIGATION_PAGE_LIST.signInPage;
        }
      }

      setIsSignedIn(checkedSignedIn);
      if (toNavigatePath !== "") {
        navigate(toNavigatePath);
      }
    }
    init();
  }, [location, setIsSignedIn]);

  return (
    <>
      <HeaderNavigation isSignedIn={isSignedIn}>
        <Outlet />
      </HeaderNavigation>
    </>
  );
}

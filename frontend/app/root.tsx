import { isRouteErrorResponse, Links, Meta, Outlet, Scripts, ScrollRestoration, useLocation, useNavigate } from "react-router";

import type { Route } from "./+types/root";
import "./app.css";
import BaseContainer from "~/components/BaseContainer";
import { getCheckSignedIn } from "~/apis/users.api";
import { useEffect, useState } from "react";
import { NAVIGATION_PAGE_LIST } from "./routes";
import { HeaderNavigation } from "./HeaderNavigation";

export const links: Route.LinksFunction = () => [
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  {
    rel: "preconnect",
    href: "https://fonts.gstatic.com",
    crossOrigin: "anonymous",
  },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  },
];

export function Layout({ children }: { children: React.ReactNode }) {
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
    <html lang='en'>
      <head>
        <meta charSet='utf-8' />
        <meta name='viewport' content='width=device-width, initial-scale=1' />
        <Meta />
        <Links />
      </head>
      <body>
        <div className='p-6'>
          <HeaderNavigation isSignedIn={isSignedIn}>
            <BaseContainer containerWidth='w-4/5'>{children}</BaseContainer>
          </HeaderNavigation>
        </div>
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return <Outlet />;
}

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  let message = "Oops!";
  let details = "An unexpected error occurred.";
  let stack: string | undefined;

  if (isRouteErrorResponse(error)) {
    message = error.status === 404 ? "404" : "Error";
    details = error.status === 404 ? "The requested page could not be found." : error.statusText || details;
  } else if (import.meta.env.DEV && error && error instanceof Error) {
    details = error.message;
    stack = error.stack;
  }

  return (
    <main className='pt-16 p-4 container mx-auto'>
      <h1>{message}</h1>
      <p>{details}</p>
      {stack && (
        <pre className='w-full p-4 overflow-x-auto'>
          <code>{stack}</code>
        </pre>
      )}
    </main>
  );
}

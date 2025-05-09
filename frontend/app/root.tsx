import { isRouteErrorResponse, Link, Links, Meta, Outlet, Scripts, ScrollRestoration } from "react-router";

import type { Route } from "./+types/root";
import "./app.css";
import BaseContainer from "~/components/BaseContainer";
import { HeaderNavigation } from "./HeaderNavigation";
import { authMiddleware } from "~/middlewares/auth-middleware";
import { AuthProvider } from "~/contexts/useAuthContext";
import { NAVIGATION_PAGE_LIST } from "./routes";

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

export const unstable_clientMiddleware = [authMiddleware];

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <AuthProvider>
      <html lang='en'>
        <head>
          <meta charSet='utf-8' />
          <meta name='viewport' content='width=device-width, initial-scale=1' />
          <Meta />
          <Links />
        </head>
        <body>
          <div>
            <HeaderNavigation>
              <BaseContainer containerWidth='w-4/5'>{children}</BaseContainer>
            </HeaderNavigation>
          </div>
          <ScrollRestoration />
          <Scripts />
        </body>
      </html>
    </AuthProvider>
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
    <>
      {import.meta.env.PROD ? (
        <div className='text-gray-800 flex items-center justify-center h-screen'>
          <div className='text-center'>
            <h1 className='text-8xl font-bold text-red-600'>500</h1>
            <h2 className='text-2xl mt-4 font-semibold'>Internal Server Error</h2>
            <p className='mt-2 text-gray-600'>サーバーでエラーが発生しました。しばらくしてから再度お試しください。</p>
            <div className='mt-6'>
              <Link to={NAVIGATION_PAGE_LIST.top} className='bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'>
                ホームに戻る
              </Link>
            </div>
          </div>
        </div>
      ) : (
        <main className='pt-16 p-4 container mx-auto'>
          <h1>{message}</h1>
          <p>{details}</p>
          {stack && (
            <pre className='w-full p-4 overflow-x-auto'>
              <code>{stack}</code>
            </pre>
          )}
        </main>
      )}
    </>
  );
}

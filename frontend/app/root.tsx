import { Link, Links, Meta, Outlet, Scripts, ScrollRestoration } from "react-router";

import { ErrorBoundary as ReactErrorBoundary } from "react-error-boundary";

import type { Route } from "./+types/root";
import "./app.css";
import BaseContainer from "~/components/BaseContainer";
import { HeaderNavigation } from "./HeaderNavigation";
import { authMiddleware } from "~/middlewares/auth-middleware";
import { AuthProvider } from "~/contexts/useAuthContext";
import { NAVIGATION_PAGE_LIST } from "./routes";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

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

function Fallback({ error }: { error: Error }) {
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
          <h1>{error.message}</h1>
          {error.stack && (
            <pre className='w-full p-4 overflow-x-auto'>
              <code>{error.stack}</code>
            </pre>
          )}
        </main>
      )}
    </>
  );
}

const queryClient = new QueryClient();

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
  return (
    <ReactErrorBoundary fallbackRender={Fallback}>
      <QueryClientProvider client={queryClient}>
        <Outlet />
      </QueryClientProvider>
    </ReactErrorBoundary>
  );
}

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  return (
    <>
      <Fallback error={error as Error} />
    </>
  );
}

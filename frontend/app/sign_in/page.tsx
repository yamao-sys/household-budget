import { SignInForm } from "~/features/sign-in/components/SignInForm";
import type { Route } from "./+types/page";
import { authContext } from "~/middlewares/auth-context";
import { useLoaderData } from "react-router";
import { useAuth } from "~/hooks/useAuth";

export async function clientLoader({ context }: Route.ClientLoaderArgs) {
  const auth = context.get(authContext);

  return { isSignedIn: !!auth?.isSignedIn };
}

export default function SignInPage() {
  const { isSignedIn } = useLoaderData<typeof clientLoader>();

  useAuth(isSignedIn);

  return (
    <>
      <SignInForm />
    </>
  );
}

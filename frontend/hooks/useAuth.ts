import { useEffect } from "react";
import { useAuthSetContext } from "~/contexts/useAuthContext";

export const useAuth = (isSignedIn: boolean, csrfToken: string) => {
  const { setAuth } = useAuthSetContext();

  useEffect(() => {
    setAuth({ isSignedIn, csrfToken });
  }, [isSignedIn, csrfToken]);

  return;
};

import { useEffect } from "react";
import { useAuthSetContext } from "~/contexts/useAuthContext";

export const useAuth = (isSignedIn: boolean) => {
  const { setAuth } = useAuthSetContext();

  useEffect(() => {
    setAuth({ isSignedIn });
  }, [isSignedIn]);

  return;
};

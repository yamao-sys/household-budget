import { createContext, type FC, useContext, useState } from "react";

type AuthSetContextType = {
  setAuth: React.Dispatch<React.SetStateAction<AuthContextType>>;
};

type AuthContextType = {
  isSignedIn: boolean;
  csrfToken: string;
};

export const AuthContext = createContext<AuthContextType>({ isSignedIn: false, csrfToken: "" });

export const AuthSetContext = createContext<AuthSetContextType>({ setAuth: () => undefined });

export const useAuthContext = () => useContext<AuthContextType>(AuthContext);

export const useAuthSetContext = () => useContext<AuthSetContextType>(AuthSetContext);

export const AuthProvider: FC<{ children: React.ReactNode }> = ({ children }) => {
  const [auth, setAuth] = useState<AuthContextType>({ isSignedIn: false, csrfToken: "" });

  return (
    <AuthContext.Provider value={auth}>
      <AuthSetContext.Provider value={{ setAuth }}>{children}</AuthSetContext.Provider>
    </AuthContext.Provider>
  );
};

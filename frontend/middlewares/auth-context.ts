import { unstable_createContext } from "react-router";

export type AuthContext = {
  isSignedIn: boolean;
};

export const authContext = unstable_createContext<AuthContext | null>();

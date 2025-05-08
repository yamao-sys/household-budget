import createClient from "openapi-fetch";
import type { paths } from "./generated/apiSchema";
import type { UserSignInInput, UserSignUpInput } from "~/types";
import { getRequestHeaders } from "./csrf.api";

const client = createClient<paths>({
  baseUrl: `${import.meta.env.VITE_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function postUserSignUp(input: UserSignUpInput, csrfToken: string) {
  const { data, response } = await client.POST("/users/signUp", {
    ...getRequestHeaders(csrfToken),
    body: input,
  });
  if (response.status === 403) {
    throw Error("Forbidden");
  }
  if (response.status === 500 || data === undefined) {
    throw Error("Internal Server Error");
  }

  return data.errors;
}

export async function postUserSignIn(input: UserSignInInput, csrfToken: string) {
  const { response } = await client.POST("/users/signIn", {
    ...getRequestHeaders(csrfToken),
    body: input,
  });
  if (response.status === 500) {
    throw Error("Internal Server Error");
  }
  if (response.status === 403) {
    throw Error("Forbidden");
  }
  if (response.status === 400) {
    return "メールアドレスまたはパスワードが正しくありません";
  }

  return "";
}

export async function getCheckSignedIn(csrfToken: string) {
  const { data, response } = await client.GET("/users/checkSignedIn", {
    ...getRequestHeaders(csrfToken),
  });
  if (data === undefined || response.status === 401) {
    return false;
  }

  return data.isSignedIn;
}

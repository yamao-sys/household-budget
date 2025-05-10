import type { UserSignInInput, UserSignUpInput, UserSignUpValidationError } from "~/types";
import { client, getRequestHeaders } from "../base/api";

export async function postUserSignUp(input: UserSignUpInput, csrfToken: string): Promise<UserSignUpValidationError> {
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

export async function postUserSignIn(input: UserSignInInput, csrfToken: string): Promise<string> {
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

export async function getCheckSignedIn(csrfToken: string): Promise<boolean> {
  const { data, response } = await client.GET("/users/checkSignedIn", {
    ...getRequestHeaders(csrfToken),
  });
  if (data === undefined || response.status === 401) {
    return false;
  }

  return data.isSignedIn;
}

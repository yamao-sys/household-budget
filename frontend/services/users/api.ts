import type { UserSignInInput, UserSignUpInput, UserSignUpValidationError } from "~/apis/model";
import { getRequestHeaders } from "../base/api";
import { getUsersCheckSignedIn, postUsersSignIn, postUsersSignUp } from "~/apis/users/users";

export async function postUserSignUp(input: UserSignUpInput, csrfToken: string): Promise<UserSignUpValidationError> {
  try {
    const res = await postUsersSignUp(input, getRequestHeaders(csrfToken));

    if (res.status === 500) {
      throw new Error(`Internal Server Error: ${res.data}`);
    }

    return res.data.errors;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

export async function postUserSignIn(input: UserSignInInput, csrfToken: string): Promise<string> {
  try {
    const res = await postUsersSignIn(input, getRequestHeaders(csrfToken));

    switch (res.status) {
      case 200:
        return "";
      case 500:
        throw new Error(`Internal Server Error: ${res.data}`);
      case 400:
        return "メールアドレスまたはパスワードが正しくありません";
    }
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

export async function getCheckSignedIn(csrfToken: string): Promise<boolean> {
  try {
    const res = await getUsersCheckSignedIn(getRequestHeaders(csrfToken));

    return res.data.isSignedIn;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

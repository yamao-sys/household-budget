import { useMutation } from "@tanstack/react-query";
import { postUserSignIn, postUserSignUp } from "./api";
import type { UserSignInInput, UserSignUpInput, UserSignUpValidationError } from "~/apis/model";

export const usePostSignUp = (
  onMutate: () => void,
  onSuccess: (data: UserSignUpValidationError) => void,
  input: UserSignUpInput,
  csrfToken: string,
) => {
  return useMutation({
    onMutate: () => onMutate,
    mutationFn: () => postUserSignUp(input, csrfToken),
    onSuccess: (data) => {
      onSuccess(data);
    },
  });
};

export const usePostSignIn = (onMutate: () => void, onSuccess: (data: string) => void, input: UserSignInInput, csrfToken: string) => {
  return useMutation({
    onMutate: () => onMutate,
    mutationFn: () => postUserSignIn(input, csrfToken),
    onSuccess: (data) => {
      onSuccess(data);
    },
  });
};

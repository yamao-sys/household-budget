import { useMutation } from "@tanstack/react-query";
import type { UserSignInInput, UserSignUpInput, UserSignUpValidationError } from "~/types";
import { postUserSignIn, postUserSignUp } from "./api";

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

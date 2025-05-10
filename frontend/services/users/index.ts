import { useMutation } from "@tanstack/react-query";
import type { UserSignUpInput, UserSignUpValidationError } from "~/types";
import { postUserSignUp } from "./api";

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

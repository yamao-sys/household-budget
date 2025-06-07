import { useCallback, useState } from "react";
import { useNavigate } from "react-router";
import type { UserSignInInput } from "~/apis/model";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";
import { usePostSignIn } from "~/services/users";

export const useSignIn = (csrfToken: string) => {
  const [userSignInInputs, setUserSignInInputs] = useState<UserSignInInput>({
    email: "",
    password: "",
  });
  const updateSignInInput = useCallback((params: Partial<UserSignInInput>) => {
    setUserSignInInputs((prev: UserSignInInput) => ({ ...prev, ...params }));
  }, []);
  const setUserSignUpTextInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      updateSignInInput({ [e.target.name]: e.target.value });
    },
    [updateSignInInput],
  );
  const [validationError, setValidationError] = useState("");

  const navigate = useNavigate();

  const initSignInValidationErrors = useCallback(() => {
    setValidationError("");
  }, []);

  const onSuccessPostSignIn = useCallback(
    (error: string) => {
      if (error !== "") {
        setValidationError(error);
        updateSignInInput({ password: "" });
        return;
      }

      window.alert("ログインしました");
      navigate(NAVIGATION_PAGE_LIST.top);
    },
    [setValidationError, updateSignInInput],
  );

  const { mutate } = usePostSignIn(initSignInValidationErrors, onSuccessPostSignIn, userSignInInputs, csrfToken);

  return {
    userSignInInputs,
    setUserSignUpTextInput,
    validationError,
    mutate,
  };
};

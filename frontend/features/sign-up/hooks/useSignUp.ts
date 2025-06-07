import { useCallback, useState } from "react";
import { useNavigate } from "react-router";
import type { UserSignUpInput, UserSignUpValidationError } from "~/apis/model";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";
import { usePostSignUp } from "~/services/users";

const INITIAL_VALIDATION_ERRORS = {
  name: [],
  email: [],
  password: [],
};

export const useSignUp = (csrfToken: string) => {
  const [userSignUpInputs, setUserSignUpInputs] = useState<UserSignUpInput>({
    name: "",
    email: "",
    password: "",
  });

  const updateSignUpInput = useCallback((params: Partial<UserSignUpInput>) => {
    setUserSignUpInputs((prev: UserSignUpInput) => ({ ...prev, ...params }));
  }, []);

  const [validationErrors, setValidationErrors] = useState<UserSignUpValidationError>(INITIAL_VALIDATION_ERRORS);

  const setSupporterSignUpTextInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      updateSignUpInput({ [e.target.name]: e.target.value });
    },
    [updateSignUpInput],
  );

  const navigate = useNavigate();

  const initSignUpValidationErrors = useCallback(() => {
    setValidationErrors(INITIAL_VALIDATION_ERRORS);
  }, []);

  const onSuccessPostSignUp = useCallback(
    (errors: UserSignUpValidationError) => {
      // バリデーションエラーがなければ、確認画面へ遷移
      if (Object.keys(errors).length === 0) {
        window.alert("会員登録が完了しました");
        navigate(NAVIGATION_PAGE_LIST.top);
        return;
      }

      // NOTE: バリデーションエラーの格納と入力パスワードのリセット
      setValidationErrors(errors);
      updateSignUpInput({ password: "" });
    },
    [setValidationErrors, updateSignUpInput],
  );

  const { mutate } = usePostSignUp(initSignUpValidationErrors, onSuccessPostSignUp, userSignUpInputs, csrfToken);

  return {
    userSignUpInputs,
    setSupporterSignUpTextInput,
    validationErrors,
    mutate,
  };
};

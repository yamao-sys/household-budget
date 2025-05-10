import type { FC } from "react";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";
import BaseButton from "~/components/BaseButton";
import BaseFormInput from "~/components/BaseFormInput";
import { useAuthContext } from "~/contexts/useAuthContext";
import { usePostSignUp } from "~/services/users";
import type { UserSignUpInput, UserSignUpValidationError } from "~/types";

const INITIAL_VALIDATION_ERRORS = {
  name: [],
  email: [],
  password: [],
};

export const SignUpForm: FC = () => {
  const [userSignUpInputs, setUserSignUpInputs] = useState<UserSignUpInput>({
    name: "",
    email: "",
    password: "",
  });

  const updateSignUpInput = useCallback((params: Partial<UserSignUpInput>) => {
    setUserSignUpInputs((prev: UserSignUpInput) => ({ ...prev, ...params }));
  }, []);

  const [validationErrors, setValidationErrors] = useState<UserSignUpValidationError>(INITIAL_VALIDATION_ERRORS);

  const { csrfToken } = useAuthContext();

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

  return (
    <>
      <h3 className='mt-16 w-full text-center text-2xl font-bold'>会員登録フォーム</h3>

      <div className='mt-8'>
        <BaseFormInput
          id='name'
          label='ユーザ名'
          name='name'
          type='text'
          value={userSignUpInputs.name}
          onChange={setSupporterSignUpTextInput}
          validationErrorMessages={validationErrors.name ?? []}
        />
      </div>

      <div className='mt-8'>
        <BaseFormInput
          id='email'
          label='Email'
          name='email'
          type='email'
          value={userSignUpInputs.email}
          onChange={setSupporterSignUpTextInput}
          validationErrorMessages={validationErrors.email ?? []}
        />
      </div>

      <div className='mt-8'>
        <BaseFormInput
          id='password'
          label='パスワード'
          name='password'
          type='password'
          value={userSignUpInputs.password}
          onChange={setSupporterSignUpTextInput}
          validationErrorMessages={validationErrors.password ?? []}
        />
      </div>

      <div className='w-full flex justify-center'>
        <div className='mt-16'>
          <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='登録する' onClick={() => mutate()} />
        </div>
      </div>
    </>
  );
};

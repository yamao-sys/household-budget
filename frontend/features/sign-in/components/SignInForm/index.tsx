import { useCallback, useState, type FC } from "react";
import { useNavigate } from "react-router";
import { postUserSignIn } from "~/apis/users.api";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";
import BaseButton from "~/components/BaseButton";
import BaseFormInput from "~/components/BaseFormInput";
import type { UserSignInInput } from "~/types";

export const SignInForm: FC = () => {
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

  const handleSignIn = useCallback(async () => {
    setValidationError("");

    const error = await postUserSignIn(userSignInInputs);

    if (error !== "") {
      setValidationError(error);
      updateSignInInput({ password: "" });
      return;
    }

    window.alert("ログインしました");
    navigate(NAVIGATION_PAGE_LIST.top);
  }, [setValidationError, userSignInInputs, updateSignInInput]);

  return (
    <>
      <h3 className='mt-16 w-full text-center text-2xl font-bold'>ログインフォーム</h3>

      {validationError && (
        <div className='w-full pt-5 text-center'>
          <p className='text-red-400'>{validationError}</p>
        </div>
      )}

      <div className='mt-8'>
        <BaseFormInput
          id='email'
          label='Email'
          name='email'
          type='email'
          value={userSignInInputs.email}
          onChange={setUserSignUpTextInput}
          validationErrorMessages={[]}
        />
      </div>

      <div className='mt-8'>
        <BaseFormInput
          id='password'
          label='パスワード'
          name='password'
          type='password'
          value={userSignInInputs.password}
          onChange={setUserSignUpTextInput}
          validationErrorMessages={[]}
        />
      </div>

      <div className='w-full flex justify-center'>
        <div className='mt-16'>
          <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='ログインする' onClick={handleSignIn} />
        </div>
      </div>
    </>
  );
};

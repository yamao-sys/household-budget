import { type FC } from "react";
import BaseButton from "~/components/BaseButton";
import BaseFormInput from "~/components/BaseFormInput";
import { useAuthContext } from "~/contexts/useAuthContext";
import { useSignIn } from "../../hooks/useSignIn";

export const SignInForm: FC = () => {
  const { csrfToken } = useAuthContext();
  const { userSignInInputs, setUserSignUpTextInput, validationError, mutate } = useSignIn(csrfToken);

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
          <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='ログインする' onClick={() => mutate()} />
        </div>
      </div>
    </>
  );
};

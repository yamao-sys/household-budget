import type { JSX } from "react";
import { memo } from "react";

type Props = {
  label: string;
  id: string;
  validationErrorMessages: string[];
} & JSX.IntrinsicElements["input"];

const BaseFormInput = memo(function BaseFormInput({ label, id, validationErrorMessages, ...props }: Props) {
  return (
    <>
      <label htmlFor={id} className='block mb-2 text-sm font-medium text-gray-900 dark:text-white text-left'>
        <span className='font-bold'>{label}</span>
      </label>
      <input
        id={id}
        className='bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
        {...props}
      />
      {validationErrorMessages.length > 0 && (
        <div className='w-full pt-5 text-left'>
          {validationErrorMessages.map((message, i) => (
            <p key={i} className='text-red-400'>
              {message}
            </p>
          ))}
        </div>
      )}
    </>
  );
});

export default BaseFormInput;

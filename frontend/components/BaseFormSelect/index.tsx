import type { JSX } from "react";
import { memo } from "react";

type Props = {
  label: string;
  id: string;
  options: [string, string][];
  validationErrorMessages: string[];
  castNumber?: boolean;
} & JSX.IntrinsicElements["select"];

const BaseFormSelect = memo(function BaseFormSelect({ label, id, options, validationErrorMessages, castNumber, ...props }: Props) {
  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    if (props.onChange === undefined) {
      return;
    }

    let value: string | number = e.target.value;
    if (castNumber) {
      value = Number(value);
    }
    let name = e.target.name;
    props.onChange({
      ...e,
      target: {
        ...e.target,
        value,
        name,
      },
    } as React.ChangeEvent<HTMLSelectElement>);
  };
  return (
    <>
      <label htmlFor={id} className='block text-sm font-medium text-gray-700 mb-1'>
        {label}
      </label>
      <select
        id={id}
        className='block w-full rounded-md border border-gray-300 p-2 shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200 focus:ring-opacity-50'
        {...props}
        onChange={handleChange}
      >
        <option value=''>-- 選択してください --</option>
        {options.map(([key, label]) => (
          <option key={key} value={key}>
            {label}
          </option>
        ))}
      </select>
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

export default BaseFormSelect;

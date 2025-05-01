import { faTimes } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import type React from "react";
import BaseButton from "~/components/BaseButton";
import BaseFormInput from "~/components/BaseFormInput";
import BaseFormSelect from "~/components/BaseFormSelect";
import { EXPENSE_CATEGORY } from "~/const/expense";
import type { Expense, StoreExpenseInput, StoreExpenseValidationError } from "~/types";

type Props = {
  inView: boolean;
  setInView: (inView: boolean) => void;
  date: string;
  expenses: Expense[];
  storeExpenseInput: StoreExpenseInput;
  setStoreExpenseTextInput: (e: React.ChangeEvent<HTMLInputElement>) => void;
  setStoreExpenseSelectInput: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  handleCreateExpense: () => Promise<void>;
  validationErrors: StoreExpenseValidationError;
};

export const DailyBudgetDialog: React.FC<Props> = ({
  inView,
  setInView,
  date,
  expenses,
  storeExpenseInput,
  setStoreExpenseTextInput,
  setStoreExpenseSelectInput,
  handleCreateExpense,
  validationErrors,
}: Props) => {
  return (
    <div
      role='dialog'
      aria-modal='true'
      className={inView ? "opacity-100 visible fixed top-1/8 left-1/4 font-bold bg-white w-1/2 flex justify-center items-center z-50" : "hidden"}
    >
      <div className='bg-white p-4 rounded shadow-lg w-full max-h-[90vh] overflow-hidden flex flex-col relative'>
        <button
          onClick={() => setInView(false)}
          className='absolute top-3 right-3 text-gray-500 hover:text-gray-700 text-xl border px-1'
          aria-label='閉じる'
        >
          <FontAwesomeIcon icon={faTimes} />
        </button>
        <h2 className='text-xl text-center font-bold mb-4'>{date} の支出</h2>

        {/* 一覧表示 */}
        <div className='overflow-y-auto mb-4 border rounded max-h-50'>
          <table className='w-full table-fixed border border-gray-300 text-sm mb-4'>
            <thead className='sticky top-0 bg-white z-10 shadow-[0px_1px_0px_0px_rgba(209,213,219,1)]'>
              <tr>
                <th className='w-1/4 text-left py-2 px-2 border border-gray-300'>金額</th>
                <th className='w-2/4 text-left py-2 px-2 border border-gray-300'>適用</th>
                <th className='w-1/4 text-left py-2 px-2 border border-gray-300'>カテゴリ</th>
              </tr>
            </thead>
            <tbody>
              {expenses.map((expense, idx) => (
                <tr key={idx}>
                  <td className='w-1/4 py-2 px-2 border border-gray-300'>¥{expense.amount}</td>
                  <td className='w-2/4 py-2 px-2 border border-gray-300 break-words'>{expense.description}</td>
                  <td className='w-1/4 py-2 px-2 border border-gray-300 break-words'>{EXPENSE_CATEGORY[expense.category]}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {/* 入力フォーム */}
        <div className='flex flex-col gap-2 mb-4 w-full'>
          <div className='mb-2'>
            <BaseFormInput
              id='amount'
              label='金額'
              name='amount'
              type='number'
              value={storeExpenseInput.amount || ""}
              onChange={setStoreExpenseTextInput}
              validationErrorMessages={validationErrors.amount ?? []}
            />
          </div>
          <div className='mb-2'>
            <BaseFormInput
              id='description'
              label='適用'
              name='description'
              type='text'
              value={storeExpenseInput.description}
              onChange={setStoreExpenseTextInput}
              validationErrorMessages={validationErrors.description ?? []}
            />
          </div>
          <div className='mb-2'>
            <BaseFormSelect
              id='category'
              label='カテゴリ'
              name='category'
              options={Object.entries(EXPENSE_CATEGORY)}
              value={storeExpenseInput.category}
              onChange={setStoreExpenseSelectInput}
              validationErrorMessages={validationErrors.category ?? []}
            />
          </div>

          <div className='w-full flex justify-center'>
            <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='登録する' onClick={handleCreateExpense} />
          </div>
        </div>
      </div>
    </div>
  );
};

import { faTimes } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import type React from "react";
import BaseButton from "~/components/BaseButton";
import BaseFormInput from "~/components/BaseFormInput";
import BaseFormSelect from "~/components/BaseFormSelect";
import { EXPENSE_CATEGORY } from "~/const/expense";
import type { Expense, Income, StoreExpenseInput, StoreExpenseValidationError } from "~/types";

type Props = {
  inView: boolean;
  setInView: (inView: boolean) => void;
  date: string;
  incomes: Income[];
  expenses: Expense[];
  storeExpenseInput: StoreExpenseInput;
  setStoreExpenseTextInput: (e: React.ChangeEvent<HTMLInputElement>) => void;
  setStoreExpenseSelectInput: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  handleCreateExpense: () => Promise<void>;
  expenseValidationErrors: StoreExpenseValidationError;
};

export const DailyBudgetDialog: React.FC<Props> = ({
  inView,
  setInView,
  date,
  incomes,
  expenses,
  storeExpenseInput,
  setStoreExpenseTextInput,
  setStoreExpenseSelectInput,
  handleCreateExpense,
  expenseValidationErrors,
}: Props) => {
  return (
    <div
      role='dialog'
      aria-modal='true'
      className={inView ? "fixed inset-0 z-50 bg-opacity-100 flex justify-center items-start overflow-y-auto py-10" : "hidden"}
    >
      <div className='bg-white p-4 rounded shadow-lg w-11/12 max-w-2xl relative overflow-y-auto max-h-[90vh]'>
        <button
          onClick={() => setInView(false)}
          className='absolute top-3 right-3 text-gray-500 hover:text-gray-700 text-xl border px-1'
          aria-label='閉じる'
        >
          <FontAwesomeIcon icon={faTimes} />
        </button>
        <h2 className='text-xl text-center font-bold mb-4'>{date} の収支</h2>

        {/* 支出 一覧表示 */}
        <h3 className='text-center font-bold mb-4'>支出</h3>
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

        <h3 className='text-center font-bold mb-4'>収入</h3>
        <div className='overflow-y-auto mb-4 border rounded max-h-50'>
          <table className='w-full table-fixed border border-gray-300 text-sm mb-4'>
            <thead className='sticky top-0 bg-white z-10 shadow-[0px_1px_0px_0px_rgba(209,213,219,1)]'>
              <tr>
                <th className='w-1/4 text-left py-2 px-2 border border-gray-300'>金額</th>
                <th className='w-2/4 text-left py-2 px-2 border border-gray-300'>顧客名</th>
              </tr>
            </thead>
            <tbody>
              {incomes.map((income, idx) => (
                <tr key={idx}>
                  <td className='w-1/4 py-2 px-2 border border-gray-300'>¥{income.amount}</td>
                  <td className='w-2/4 py-2 px-2 border border-gray-300 break-words'>{income.clientName}</td>
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
              validationErrorMessages={expenseValidationErrors.amount ?? []}
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
              validationErrorMessages={expenseValidationErrors.description ?? []}
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
              validationErrorMessages={expenseValidationErrors.category ?? []}
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

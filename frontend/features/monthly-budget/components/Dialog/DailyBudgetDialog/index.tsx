import { faTimes } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useQueryClient } from "@tanstack/react-query";
import type React from "react";
import BaseButton from "~/components/BaseButton";
import BaseFormInput from "~/components/BaseFormInput";
import BaseFormSelect from "~/components/BaseFormSelect";
import { EXPENSE_CATEGORY } from "~/const/expense";
import { useAuthContext } from "~/contexts/useAuthContext";
import { useGetExpenses, usePostCreateExpense } from "~/services/expenses";
import { useGetIncomes, usePostCreateIncome } from "~/services/incomes";
import type {
  StoreExpenseInput,
  StoreExpenseResponse,
  StoreExpenseValidationError,
  StoreIncomeInput,
  StoreIncomeResponse,
  StoreIncomeValidationError,
} from "~/types";

type Props = {
  inView: boolean;
  setInView: (inView: boolean) => void;
  date: string;
  // NOTE: 支出登録関連
  storeExpenseInput: StoreExpenseInput;
  setStoreExpenseTextInput: (e: React.ChangeEvent<HTMLInputElement>) => void;
  setStoreExpenseSelectInput: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  onPostCreateExpenseMutate: () => void;
  onPostCreateExpenseSuccess: (data: StoreExpenseResponse) => void;
  expenseValidationErrors: StoreExpenseValidationError;
  // NOTE: 収入登録関連
  storeIncomeInput: StoreIncomeInput;
  setStoreIncomeTextInput: (e: React.ChangeEvent<HTMLInputElement>) => void;
  onPostCreateIncomeMutate: () => void;
  onPostCreateIncomeSuccess: (data: StoreIncomeResponse) => void;
  incomeValidationErrors: StoreIncomeValidationError;
};

export const DailyBudgetDialog: React.FC<Props> = ({
  inView,
  setInView,
  date,
  // NOTE: 支出登録関連
  storeExpenseInput,
  setStoreExpenseTextInput,
  setStoreExpenseSelectInput,
  onPostCreateExpenseMutate,
  onPostCreateExpenseSuccess,
  expenseValidationErrors,
  // NOTE: 収入登録関連
  storeIncomeInput,
  setStoreIncomeTextInput,
  onPostCreateIncomeMutate,
  onPostCreateIncomeSuccess,
  incomeValidationErrors,
}: Props) => {
  const { csrfToken } = useAuthContext();

  const {
    data: selectedDateExpenses,
    isPending: isSelectedDateExpensesPending,
    isError: isSelectedDateExpensesError,
  } = useGetExpenses(date, date, csrfToken);

  const {
    data: selectedDateIncomes,
    isPending: isSelectedDateIncomesPending,
    isError: isSelectedDateIncomesError,
  } = useGetIncomes(date, date, csrfToken);

  const queryClient = useQueryClient();

  const {
    mutate: poseCreateExpenseMutate,
    // isPending: isPostCreateExpensePending,
    // isError: isPostCreateExpenseError,
  } = usePostCreateExpense(queryClient, onPostCreateExpenseMutate, onPostCreateExpenseSuccess, storeExpenseInput, date, csrfToken);

  const {
    mutate: poseCreateIncomeMutate,
    // isPending: isPostCreateIncomePending,
    // isError: isPostCreateIncomeError,
  } = usePostCreateIncome(queryClient, onPostCreateIncomeMutate, onPostCreateIncomeSuccess, storeIncomeInput, date, csrfToken);

  return (
    <div
      role='dialog'
      aria-modal='true'
      className={inView ? "fixed inset-0 z-150 bg-opacity-100 flex justify-center items-start overflow-y-auto py-10" : "hidden"}
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
        {isSelectedDateExpensesPending && <div className='text-center'>Loading...</div>}
        {isSelectedDateExpensesError && <div className='text-center'>Error loading expenses</div>}

        {!!selectedDateExpenses && !isSelectedDateExpensesPending && !isSelectedDateExpensesError && (
          <>
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
                  {selectedDateExpenses.map((expense, idx) => (
                    <tr key={idx}>
                      <td className='w-1/4 py-2 px-2 border border-gray-300'>¥{expense.amount.toLocaleString()}</td>
                      <td className='w-2/4 py-2 px-2 border border-gray-300 break-words'>{expense.description}</td>
                      <td className='w-1/4 py-2 px-2 border border-gray-300 break-words'>{EXPENSE_CATEGORY[expense.category]}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </>
        )}

        {/* 収入 一覧表示 */}
        {isSelectedDateIncomesPending && <div className='text-center'>Loading...</div>}
        {isSelectedDateIncomesError && <div className='text-center'>Error loading expenses</div>}

        {!!selectedDateIncomes && !isSelectedDateIncomesPending && !isSelectedDateIncomesError && (
          <>
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
                  {selectedDateIncomes.map((income, idx) => (
                    <tr key={idx}>
                      <td className='w-1/4 py-2 px-2 border border-gray-300'>¥{income.amount.toLocaleString()}</td>
                      <td className='w-2/4 py-2 px-2 border border-gray-300 break-words'>{income.clientName}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </>
        )}

        {/* 入力フォーム */}
        <div className='flex flex-col gap-2 p-4 mb-4 w-full border rounded'>
          <h3 className='text-center font-bold mb-4'>支出登録</h3>
          <div className='w-1/2 mx-auto mb-2'>
            <BaseFormInput
              id='expense-amount'
              label='支出金額'
              name='amount'
              type='number'
              value={storeExpenseInput.amount || ""}
              onChange={setStoreExpenseTextInput}
              validationErrorMessages={expenseValidationErrors.amount ?? []}
            />
          </div>
          <div className='w-1/2 mx-auto mb-2'>
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
          <div className='w-1/2 mx-auto mb-2'>
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
            <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='支出を登録する' onClick={() => poseCreateExpenseMutate()} />
          </div>
        </div>

        <div className='flex flex-col gap-2 p-4 mb-4 w-full border rounded'>
          <h3 className='text-center font-bold mb-4'>収入登録</h3>
          <div className='w-1/2 mx-auto mb-2'>
            <BaseFormInput
              id='income-amount'
              label='収入金額'
              name='amount'
              type='number'
              value={storeIncomeInput.amount || ""}
              onChange={setStoreIncomeTextInput}
              validationErrorMessages={incomeValidationErrors.amount ?? []}
            />
          </div>
          <div className='w-1/2 mx-auto mb-2'>
            <BaseFormInput
              id='client-name'
              label='顧客名'
              name='clientName'
              type='text'
              value={storeIncomeInput.clientName}
              onChange={setStoreIncomeTextInput}
              validationErrorMessages={incomeValidationErrors.clientName ?? []}
            />
          </div>

          <div className='w-full flex justify-center'>
            <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='収入を登録する' onClick={() => poseCreateIncomeMutate()} />
          </div>
        </div>
      </div>
    </div>
  );
};

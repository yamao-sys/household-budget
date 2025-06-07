import { useMemo } from "react";
import { EXPENSE_CATEGORY } from "~/const/expense";
import { getDateString, getMonthLocaleString } from "~/lib/date";
import { useGetExpenseCategoryTotalAmounts } from "~/services/expenses";
import { useGetIncomeClientTotalAmounts } from "~/services/incomes";

type Props = {
  monthDate: Date;
  csrfToken: string;
};

export const MonthlyBudgetDetail: React.FC<Props> = ({ monthDate, csrfToken }: Props) => {
  const monthBeginningDate = new Date(monthDate.getFullYear(), monthDate.getMonth(), 1);
  const monthEndDate = new Date(monthDate.getFullYear(), monthDate.getMonth() + 1, 0);

  const {
    data: expenseCategoryTotalAmounts,
    isPending: isGetExpenseCategoryTotalAmountsPending,
    isError: isGetExpenseCategoryTotalAmountsError,
  } = useGetExpenseCategoryTotalAmounts(getDateString(monthBeginningDate), getDateString(monthEndDate), csrfToken);

  const {
    data: incomeClientTotalAmounts,
    isPending: isGetIncomeClientTotalAmountsPending,
    isError: isGetIncomeClientTotalAmountsError,
  } = useGetIncomeClientTotalAmounts(getDateString(monthBeginningDate), getDateString(monthEndDate), csrfToken);

  const expenseTotalAmounts = useMemo(
    () => (!!expenseCategoryTotalAmounts ? expenseCategoryTotalAmounts.reduce((acc, totalAmounts) => acc + totalAmounts.totalAmount, 0) : 0),
    [expenseCategoryTotalAmounts],
  );
  const incomeTotalAmounts = useMemo(
    () => (!!incomeClientTotalAmounts ? incomeClientTotalAmounts.reduce((acc, totalAmounts) => acc + totalAmounts.totalAmount, 0) : 0),
    [incomeClientTotalAmounts],
  );
  const balance = useMemo(() => incomeTotalAmounts - expenseTotalAmounts, [incomeTotalAmounts, expenseTotalAmounts]);

  return (
    <div className='mx-auto mt-4'>
      <h2 className='text-2xl text-center font-bold mb-8'>{getMonthLocaleString(monthDate)}の収支詳細</h2>

      {isGetExpenseCategoryTotalAmountsPending || isGetIncomeClientTotalAmountsPending ? (
        <div className='text-center'>
          <p className='text-gray-500'>Loading...</p>
        </div>
      ) : isGetExpenseCategoryTotalAmountsError || isGetIncomeClientTotalAmountsError ? (
        <div className='text-center'>
          <p className='text-red-500'>Error occurred while fetching data.</p>
        </div>
      ) : (
        <>
          <div className='mb-8'>
            <div className='flex justify-center'>
              <div className='text-green-700 mr-4'>{`収入合計: ¥${incomeTotalAmounts.toLocaleString()}`}</div>
              <div className='text-red-700 mr-4'>{`支出合計: ¥${expenseTotalAmounts.toLocaleString()}`}</div>
              <div className='text-blue-700 font-bold'>{`利益: ¥${balance.toLocaleString()}`}</div>
            </div>
          </div>

          <div className='mb-8'>
            <h3 className='text-xl text-center font-bold mb-4'>支出</h3>
            <table className='w-full table-fixed border border-gray-300 text-sm mb-4'>
              <thead className='sticky top-0 bg-white z-10 shadow-[0px_1px_0px_0px_rgba(209,213,219,1)]'>
                <tr>
                  <th className='w-1/2 text-left py-2 px-2 border border-gray-300'>カテゴリ</th>
                  <th className='w-1/2 text-left py-2 px-2 border border-gray-300'>支出合計額</th>
                </tr>
              </thead>
              <tbody>
                {Object.entries(EXPENSE_CATEGORY).map(([categoryIdx, label], idx) => (
                  <tr key={idx}>
                    <td className='w-1/2 py-2 px-2 border border-gray-300'>{label}</td>
                    <td className='w-1/2 py-2 px-2 border border-gray-300'>
                      ¥
                      {(
                        expenseCategoryTotalAmounts?.find((totalAmounts) => totalAmounts.category === Number(categoryIdx))?.totalAmount ?? 0
                      ).toLocaleString()}
                    </td>
                  </tr>
                ))}
                <tr>
                  <td className='w-1/2 py-2 px-2 border border-gray-300 font-bold'>合計</td>
                  <td className='w-1/2 py-2 px-2 border border-gray-300 font-bold'>¥{expenseTotalAmounts.toLocaleString()}</td>
                </tr>
              </tbody>
            </table>
          </div>

          <div className='mb-8'>
            <h3 className='text-xl text-center font-bold mb-4'>収入</h3>
            <table className='w-full table-fixed border border-gray-300 text-sm mb-4'>
              <thead className='sticky top-0 bg-white z-10 shadow-[0px_1px_0px_0px_rgba(209,213,219,1)]'>
                <tr>
                  <th className='w-1/2 text-left py-2 px-2 border border-gray-300'>顧客名</th>
                  <th className='w-1/2 text-left py-2 px-2 border border-gray-300'>支出合計額</th>
                </tr>
              </thead>
              <tbody>
                {incomeClientTotalAmounts?.map(({ clientName, totalAmount }, idx) => (
                  <tr key={idx}>
                    <td className='w-1/2 py-2 px-2 border border-gray-300'>{clientName}</td>
                    <td className='w-1/2 py-2 px-2 border border-gray-300'>¥{totalAmount.toLocaleString()}</td>
                  </tr>
                ))}
                <tr>
                  <td className='w-1/2 py-2 px-2 border border-gray-300 font-bold'>合計</td>
                  <td className='w-1/2 py-2 px-2 border border-gray-300 font-bold'>¥{incomeTotalAmounts.toLocaleString()}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </>
      )}
    </div>
  );
};

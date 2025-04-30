import { EXPENSE_CATEGORY } from "~/const/expense";
import { getMonthLocaleString } from "~/lib/date";
import type { CategoryTotalAmountLists } from "~/types";

type Props = {
  monthDate: Date;
  categoryTotalAmounts: CategoryTotalAmountLists;
};

export const MonthlyBudgetDetail: React.FC<Props> = ({ monthDate, categoryTotalAmounts }: Props) => {
  return (
    <div className='mx-auto mt-4'>
      <h2 className='text-2xl text-center font-bold mb-4'>{getMonthLocaleString(monthDate)}の支出詳細</h2>

      <div className='mb-4'>
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
                  ¥{categoryTotalAmounts.find((totalAmounts) => totalAmounts.category === Number(categoryIdx))?.totalAmount ?? 0}
                </td>
              </tr>
            ))}
            <tr>
              <td className='w-1/2 py-2 px-2 border border-gray-300 font-bold'>合計</td>
              <td className='w-1/2 py-2 px-2 border border-gray-300 font-bold'>
                ¥{categoryTotalAmounts.reduce((acc, totalAmounts) => acc + totalAmounts.totalAmount, 0)}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

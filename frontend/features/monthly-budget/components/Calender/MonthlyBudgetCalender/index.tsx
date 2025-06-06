import React, { useEffect } from "react";
// NOTE: FullCalendarコンポーネント。
import FullCalendar from "@fullcalendar/react";
// NOTE: FullCalendarで月表示を可能にするモジュール。
import dayGridPlugin from "@fullcalendar/daygrid";
// NOTE: FullCalendarで日付や時間が選択できるようになるモジュール。
import interactionPlugin from "@fullcalendar/interaction";

import { Link } from "react-router";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";
import { getDateString, getMonthString } from "~/lib/date";
import { useMonthlyBudgetCalender } from "~/features/monthly-budget/hooks/useMonthlyBudgetCalender";
import { DailyBudgetDialog } from "../../Dialog/DailyBudgetDialog";
import { useGetExpenseTotalAmounts } from "~/services/expenses";
import { useAuthContext } from "~/contexts/useAuthContext";
import { useGetIncomeTotalAmounts } from "~/services/incomes";

export const MonthlyBudgetCalender: React.FC = () => {
  /**
   * 予定を追加する際にCalendarオブジェクトのメソッドを使用する必要がある。
   * (CalendarオブジェクトはRef経由でアクセスする必要がある。)
   */
  const ref = React.createRef<any>();

  const {
    selectedMonth,
    summary,
    handleDatesSet,
    handleDateClick,
    setEvents,
    events,

    dialog,
  } = useMonthlyBudgetCalender();

  const { csrfToken } = useAuthContext();

  const {
    data: expenseTotalAmounts,
    isPending: isGetExpenseTotalAmountsPending,
    isError: isGetExpenseTotalAmountsError,
  } = useGetExpenseTotalAmounts(getDateString(selectedMonth.beginning), getDateString(selectedMonth.end), csrfToken);

  const {
    data: incomeTotalIncomes,
    isPending: isGetIncomeTotalIncomesPending,
    isError: isGetIncomeTotalIncomesError,
  } = useGetIncomeTotalAmounts(getDateString(selectedMonth.beginning), getDateString(selectedMonth.end), csrfToken);

  useEffect(() => {
    if (expenseTotalAmounts === undefined || incomeTotalIncomes === undefined) return;

    setEvents([...(expenseTotalAmounts ?? []), ...(incomeTotalIncomes ?? [])]);
  }, [expenseTotalAmounts, incomeTotalIncomes]);

  return (
    <div className='mx-auto mt-4'>
      {isGetExpenseTotalAmountsPending || isGetIncomeTotalIncomesPending ? (
        <div className='text-center'>
          <p className='text-gray-500'>Loading...</p>
        </div>
      ) : isGetExpenseTotalAmountsError || isGetIncomeTotalIncomesError ? (
        <div className='text-center'>
          <p className='text-red-500'>Error occurred while fetching data.</p>
        </div>
      ) : (
        <>
          {/* 選択した日付の支出と登録フォーム */}
          <DailyBudgetDialog
            inView={dialog.inView}
            setInView={dialog.setInView}
            date={dialog.selectedDate}
            storeExpenseInput={dialog.store.storeExpenseInput}
            setStoreExpenseTextInput={dialog.store.setStoreExpenseTextInput}
            setStoreExpenseSelectInput={dialog.store.setStoreExpenseSelectInput}
            onPostCreateExpenseMutate={dialog.store.initExpenseValidationErrors}
            onPostCreateExpenseSuccess={dialog.store.onSuccessPostCreateExpense}
            expenseValidationErrors={dialog.store.expenseValidationErrors}
            storeIncomeInput={dialog.store.storeIncomeInput}
            setStoreIncomeTextInput={dialog.store.setStoreIncomeTextInput}
            onPostCreateIncomeMutate={dialog.store.initIncomeValidationErrors}
            onPostCreateIncomeSuccess={dialog.store.onSuccessPostCreateIncome}
            incomeValidationErrors={dialog.store.incomeValidationErrors}
          />

          {/* 合計表示 */}
          <div className='mb-4 p-4 bg-gray-100 rounded-lg shadow text-sm'>
            <div className='flex'>
              <div className='text-green-700 mr-4'>{`収入合計: ¥${summary.totalIncome.toLocaleString()}`}</div>
              <div className='text-red-700 mr-4'>{`支出合計: ¥${summary.totalExpense.toLocaleString()}`}</div>
              <div className='text-blue-700 underline'>
                <Link to={`${NAVIGATION_PAGE_LIST.monthlyBudgetPage}/${getMonthString(selectedMonth.beginning)}`}>収支詳細へ</Link>
              </div>
            </div>
            <div className='mt-2 font-bold'>{`利益: ¥${(summary.totalIncome - summary.totalExpense).toLocaleString()}`}</div>
          </div>
        </>
      )}

      <FullCalendar
        locale='ja'
        plugins={[dayGridPlugin, interactionPlugin]}
        initialView='dayGridMonth' // NOTE: カレンダーの初期表示設定
        events={events}
        eventContent={(arg) => {
          const { extendProps } = arg.event.extendedProps;
          return (
            <div className='text-xs bg-white p-1 rounded shadow-md'>
              <span className={extendProps.type === "income" ? "text-green-600 font-bold" : "text-red-600"}>
                {`${extendProps.type === "income" ? "収入: " : "支出: "}¥${extendProps.totalAmount.toLocaleString()}`}
              </span>
            </div>
          );
        }}
        selectable={true} // NOTE: 日付選択を可能にする。interactionPluginが有効になっている場合のみ。
        businessHours={{
          daysOfWeek: [0, 1, 2, 3, 4, 5, 6], // NOTE: 0:日曜 〜 6:土曜
          startTime: "00:00",
          endTIme: "24:00",
        }}
        weekends={true} // NOTE: 週末を強調表示する。
        titleFormat={{
          year: "numeric",
          month: "short",
        }}
        headerToolbar={{
          start: "title",
          center: "prev, next, today",
          end: "dayGridMonth",
        }}
        ref={ref}
        dateClick={handleDateClick}
        height='auto'
        datesSet={handleDatesSet}
      />
    </div>
  );
};

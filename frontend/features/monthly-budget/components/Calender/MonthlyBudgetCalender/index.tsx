import React from "react";
// NOTE: FullCalendarコンポーネント。
import FullCalendar from "@fullcalendar/react";
// NOTE: FullCalendarで月表示を可能にするモジュール。
import dayGridPlugin from "@fullcalendar/daygrid";
// NOTE: FullCalendarで日付や時間が選択できるようになるモジュール。
import interactionPlugin from "@fullcalendar/interaction";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faTimes } from "@fortawesome/free-solid-svg-icons";
import BaseFormInput from "~/components/BaseFormInput";
import BaseButton from "~/components/BaseButton";
import { EXPENSE_CATEGORY } from "~/const/expense";
import BaseFormSelect from "~/components/BaseFormSelect";
import { Link } from "react-router";
import { NAVIGATION_PAGE_LIST } from "~/app/routes";
import { getMonthString } from "~/lib/date";
import { useMonthlyBudgetCalender } from "~/features/monthly-budget/hooks/useMonthlyBudgetCalender";

export const MonthlyBudgetCalender: React.FC = () => {
  /**
   * 予定を追加する際にCalendarオブジェクトのメソッドを使用する必要がある。
   * (CalendarオブジェクトはRef経由でアクセスする必要がある。)
   */
  const ref = React.createRef<any>();

  const {
    currentMonthDate,
    summary,
    handleDatesSet,
    handleDateClick,
    events,

    dialog,
  } = useMonthlyBudgetCalender();

  const formElement = (
    <div
      role='dialog'
      aria-modal='true'
      className={
        dialog.inView ? "opacity-100 visible fixed top-1/8 left-1/4 font-bold bg-white w-1/2 flex justify-center items-center z-50" : "hidden"
      }
    >
      <div className='bg-white p-4 rounded shadow-lg w-full max-h-[90vh] overflow-hidden flex flex-col relative'>
        <button
          onClick={() => dialog.setInView(false)}
          className='absolute top-3 right-3 text-gray-500 hover:text-gray-700 text-xl border px-1'
          aria-label='閉じる'
        >
          <FontAwesomeIcon icon={faTimes} />
        </button>
        <h2 className='text-xl text-center font-bold mb-4'>{dialog.selectedDate} の支出</h2>

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
              {dialog.selectedDateExpenses.map((expense, idx) => (
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
              value={dialog.store.storeExpenseInput.amount || ""}
              onChange={dialog.store.setStoreExpenseTextInput}
              validationErrorMessages={dialog.store.validationErrors.amount ?? []}
            />
          </div>
          <div className='mb-2'>
            <BaseFormInput
              id='description'
              label='適用'
              name='description'
              type='text'
              value={dialog.store.storeExpenseInput.description}
              onChange={dialog.store.setStoreExpenseTextInput}
              validationErrorMessages={dialog.store.validationErrors.description ?? []}
            />
          </div>
          <div className='mb-2'>
            <BaseFormSelect
              id='category'
              label='カテゴリ'
              name='category'
              options={Object.entries(EXPENSE_CATEGORY)}
              value={dialog.store.storeExpenseInput.category}
              onChange={dialog.store.setStoreExpenseSelectInput}
              validationErrorMessages={dialog.store.validationErrors.category ?? []}
            />
          </div>

          <div className='w-full flex justify-center'>
            <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='登録する' onClick={dialog.store.handleCreateExpense} />
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <div className='mx-auto mt-4'>
      {formElement}
      {/* 合計表示 */}
      <div className='mb-4 p-4 bg-gray-100 rounded-lg shadow text-sm'>
        <div className='flex'>
          {/* <div className="text-green-700">収入合計: ¥{summary.totalIncome.toLocaleString()}</div> */}
          <div className='text-red-700 mr-4'>{`支出合計: ¥${summary.totalExpense.toLocaleString()}`}</div>
          <div className='text-blue-700 underline'>
            <Link to={`${NAVIGATION_PAGE_LIST.monthlyBudgetPage}/${getMonthString(currentMonthDate)}`}>支出詳細へ</Link>
          </div>
        </div>
        {/* <div className="mt-2 font-bold">
          残高: ¥{(summary.totalIncome - summary.totalExpense).toLocaleString()}
        </div> */}
      </div>

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
                {`${extendProps.type === "income" ? "収入: " : "支出: "}¥${extendProps.totalAmount}`}
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

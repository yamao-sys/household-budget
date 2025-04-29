import React, { useState } from "react";
// NOTE: FullCalendarコンポーネント。
import FullCalendar from "@fullcalendar/react";
// NOTE: FullCalendarで月表示を可能にするモジュール。
import dayGridPlugin from "@fullcalendar/daygrid";
// NOTE: FullCalendarで日付や時間が選択できるようになるモジュール。
import interactionPlugin, { type DateClickArg } from "@fullcalendar/interaction";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import type { DatesSetArg } from "@fullcalendar/core";
import { faTimes } from "@fortawesome/free-solid-svg-icons";
import type { TotalAmountLists } from "~/types";
import { getTotalAmounts } from "~/apis/expenses.api";

export const MonthlyBudgetCalender: React.FC = () => {
  /**
   * 予定を追加する際にCalendarオブジェクトのメソッドを使用する必要がある。
   * (CalendarオブジェクトはRef経由でアクセスする必要がある。)
   */
  const ref = React.createRef<any>();

  const [events, setEvents] = useState<TotalAmountLists>([]);
  const [inView, setInView] = useState(false);
  const [selectedDate, setSelectedDate] = useState("");
  const [currentMonthDate, setCurrentMonthDate] = useState<Date>(new Date());

  const handleDatesSet = async (arg: DatesSetArg) => {
    const selectedMonthBeginningDate = arg.view.currentStart;
    const currentEnd = arg.view.currentEnd;
    currentEnd.setDate(currentEnd.getDate() - 1);
    const selectedMonthEndDate = currentEnd;

    // TODO: ここをTanstack Query等を使用してキャッシュする
    const fetchedTotalAmounts = await getTotalAmounts(
      selectedMonthBeginningDate.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit", day: "2-digit" }).replaceAll("/", "-"),
      selectedMonthEndDate.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit", day: "2-digit" }).replaceAll("/", "-"),
    );
    setCurrentMonthDate(selectedMonthBeginningDate);
    setEvents(fetchedTotalAmounts);
  };

  const handleDateClick = (arg: DateClickArg) => {
    if (arg.date.getMonth() !== currentMonthDate.getMonth()) return;

    setSelectedDate(arg.date.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit", day: "2-digit" }).replaceAll("/", "-"));
    setInView(true);
  };
  const formElement = (
    <div className={inView ? "opacity-100 visible fixed top-1/8 left-1/4 font-bold bg-white w-1/2 flex justify-center items-center z-50" : "hidden"}>
      <div className='bg-white p-4 rounded shadow-lg w-full max-h-[90vh] overflow-hidden flex flex-col relative'>
        <button
          onClick={() => setInView(false)}
          className='absolute top-3 right-3 text-gray-500 hover:text-gray-700 text-xl border px-1'
          aria-label='閉じる'
        >
          <FontAwesomeIcon icon={faTimes} />
        </button>
        <h2 className='text-xl text-center font-bold mb-4'>{selectedDate} の支出</h2>

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
              {[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map((e, idx) => (
                <tr key={idx}>
                  <td className='w-1/4 py-2 px-2 border border-gray-300'>¥10,000</td>
                  <td className='w-2/4 py-2 px-2 border border-gray-300'>西友</td>
                  <td className='w-1/4 py-2 px-2 border border-gray-300'>日用品</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {/* 入力フォーム */}
        <form className='flex flex-col gap-2 mb-4 w-full'>
          <input name='amount' type='number' placeholder='金額' required className='border px-2 py-1 rounded' />
          <input name='usage' type='text' placeholder='適用' required className='border px-2 py-1 rounded' />
          <input name='category' type='text' placeholder='カテゴリ' required className='border px-2 py-1 rounded' />
          <button type='submit' className='mx-auto w-50 bg-blue-500 text-white px-4 py-1 rounded'>
            追加
          </button>
        </form>
      </div>
    </div>
  );

  return (
    <div className='mx-auto mt-10'>
      {formElement}
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
                {extendProps.type === "income" ? "収入: " : "支出: "}¥{extendProps.totalAmount}
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

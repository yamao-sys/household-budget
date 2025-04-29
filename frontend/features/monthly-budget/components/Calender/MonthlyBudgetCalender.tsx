import React, { useCallback, useState } from "react";
// NOTE: FullCalendarコンポーネント。
import FullCalendar from "@fullcalendar/react";
// NOTE: FullCalendarで月表示を可能にするモジュール。
import dayGridPlugin from "@fullcalendar/daygrid";
// NOTE: FullCalendarで日付や時間が選択できるようになるモジュール。
import interactionPlugin, { type DateClickArg } from "@fullcalendar/interaction";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import type { DatesSetArg } from "@fullcalendar/core";
import { faTimes } from "@fortawesome/free-solid-svg-icons";
import type { StoreExpenseInput, Expense, TotalAmountLists, StoreExpenseValidationError } from "~/types";
import { getExpenses, getTotalAmounts, postCreateExpense } from "~/apis/expenses.api";
import BaseFormInput from "~/components/BaseFormInput";
import BaseButton from "~/components/BaseButton";

const INITIAL_STORE_EXPENSE_INPUT = {
  paidAt: new Date(),
  amount: 0,
  category: 1,
  description: "",
};
const INITIAL_VALIDATION_ERRORS = {
  amount: [],
  category: [],
  description: [],
};

export const MonthlyBudgetCalender: React.FC = () => {
  /**
   * 予定を追加する際にCalendarオブジェクトのメソッドを使用する必要がある。
   * (CalendarオブジェクトはRef経由でアクセスする必要がある。)
   */
  const ref = React.createRef<any>();

  const [events, setEvents] = useState<TotalAmountLists>([]);
  const [inView, setInView] = useState(false);
  const [selectedDate, setSelectedDate] = useState("");
  const [selectedDateExpenses, setSelectedDateExpenses] = useState<Expense[]>([]);
  const [currentMonthDate, setCurrentMonthDate] = useState<Date>(new Date());

  const [storeExpenseInput, setStoreExpenseInput] = useState<StoreExpenseInput>(INITIAL_STORE_EXPENSE_INPUT);
  const updateStoreExpenseInput = useCallback((params: Partial<StoreExpenseInput>) => {
    setStoreExpenseInput((prev: StoreExpenseInput) => ({ ...prev, ...params }));
  }, []);
  const setStoreExpenseTextInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      updateStoreExpenseInput({ [e.target.name]: e.target.value });
    },
    [updateStoreExpenseInput],
  );
  const [validationErrors, setValidationErrors] = useState<StoreExpenseValidationError>(INITIAL_VALIDATION_ERRORS);

  // NOTE: 月が変更された時の処理
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

  // NOTE: 日が選択された時の処理
  const handleDateClick = async (arg: DateClickArg) => {
    // NOTE: 選択月以外の日付のクリックは無効にする
    if (arg.date.getMonth() !== currentMonthDate.getMonth()) return;

    const date = arg.date.toLocaleDateString("ja-jp", { year: "numeric", month: "2-digit", day: "2-digit" }).replaceAll("/", "-");
    setSelectedDateExpenses(await getExpenses(date, date));
    setSelectedDate(date);
    setInView(true);
    setStoreExpenseInput((prev: StoreExpenseInput) => ({ ...prev, ...{ paidAt: arg.date } }));
  };

  const handleCreateExpense = async () => {
    setValidationErrors(INITIAL_VALIDATION_ERRORS);

    const { expense, errors } = await postCreateExpense(storeExpenseInput);
    if (Object.keys(errors).length > 0) {
      setValidationErrors(errors);
      return;
    }

    window.alert("支出を登録しました");

    // NOTE: 選択月の支出に反映する
    const otherDateEvents = events.filter((event) => String(event.date) !== selectedDate);
    const currentDateEvent = events.find((event) => String(event.date) === selectedDate);
    const newTotalAmount = currentDateEvent?.extendProps.totalAmount ?? 0;
    setEvents([
      ...otherDateEvents,
      {
        date: expense.paidAt,
        extendProps: {
          type: "expense",
          totalAmount: newTotalAmount + expense.amount,
        },
      },
    ]);
    setInView(false);
    setStoreExpenseInput(INITIAL_STORE_EXPENSE_INPUT);
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
              {selectedDateExpenses.map((expense, idx) => (
                <tr key={idx}>
                  <td className='w-1/4 py-2 px-2 border border-gray-300'>¥{expense.amount}</td>
                  <td className='w-2/4 py-2 px-2 border border-gray-300 break-words'>{expense.description}</td>
                  <td className='w-1/4 py-2 px-2 border border-gray-300 break-words'>{expense.category}</td>
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

          <div className='w-full flex justify-center'>
            <div className='mt-16'>
              <BaseButton borderColor='border-green-500' bgColor='bg-green-500' label='登録する' onClick={handleCreateExpense} />
            </div>
          </div>
        </div>
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

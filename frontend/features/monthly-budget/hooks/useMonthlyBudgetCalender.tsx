import type { DatesSetArg } from "@fullcalendar/core/index.js";
import type { DateClickArg } from "@fullcalendar/interaction/index.js";
import { useCallback, useMemo, useState } from "react";
import { getExpenses, getTotalAmounts, postCreateExpense } from "~/apis/expenses.api";
import { getIncomeTotalAmounts } from "~/apis/incomes.api";
import { getDateString } from "~/lib/date";
import type { Expense, StoreExpenseInput, StoreExpenseValidationError, TotalAmountLists } from "~/types";

const INITIAL_STORE_EXPENSE_INPUT = {
  paidAt: new Date(),
  amount: 0,
  category: 0,
  description: "",
};
const INITIAL_VALIDATION_ERRORS = {
  amount: [],
  category: [],
  description: [],
};

export const useMonthlyBudgetCalender = () => {
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
  const setStoreExpenseSelectInput = useCallback(
    (e: React.ChangeEvent<HTMLSelectElement>) => {
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
    const fetchedExpenseTotalAmounts = await getTotalAmounts(getDateString(selectedMonthBeginningDate), getDateString(selectedMonthEndDate));
    const fetchedIncomeTotalAmounts = await getIncomeTotalAmounts(getDateString(selectedMonthBeginningDate), getDateString(selectedMonthEndDate));

    setCurrentMonthDate(selectedMonthBeginningDate);
    setEvents([...(fetchedExpenseTotalAmounts ?? []), ...(fetchedIncomeTotalAmounts ?? [])]);
  };

  // NOTE: 日が選択された時の処理
  const handleDateClick = async (arg: DateClickArg) => {
    // NOTE: 選択月以外の日付のクリックは無効にする
    if (arg.date.getMonth() !== currentMonthDate.getMonth()) return;

    const date = getDateString(arg.date);
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

  // NOTE: 表示中の月の収支集計
  const summary = useMemo(() => {
    let totalIncome = 0;
    let totalExpense = 0;
    if (!events) return { totalIncome, totalExpense };

    for (const e of events) {
      if (e.extendProps.type === "income") {
        totalIncome += e.extendProps.totalAmount;
      } else if (e.extendProps.type === "expense") {
        totalExpense += e.extendProps.totalAmount;
      }
    }
    return { totalIncome, totalExpense };
  }, [events]);

  return {
    currentMonthDate,
    summary,
    handleDatesSet,
    handleDateClick,
    events,

    dialog: {
      inView,
      setInView,
      selectedDate,
      selectedDateExpenses,
      store: {
        handleCreateExpense,
        storeExpenseInput,
        validationErrors,
        setStoreExpenseTextInput,
        setStoreExpenseSelectInput,
      },
    },
  };
};

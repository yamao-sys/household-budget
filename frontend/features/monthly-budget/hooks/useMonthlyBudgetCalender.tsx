import type { DatesSetArg } from "@fullcalendar/core/index.js";
import type { DateClickArg } from "@fullcalendar/interaction/index.js";
import { useCallback, useMemo, useState } from "react";
import { getDateString } from "~/lib/date";
import type {
  StoreExpenseInput,
  StoreExpenseResponse,
  StoreExpenseValidationError,
  StoreIncomeInput,
  StoreIncomeResponse,
  StoreIncomeValidationError,
  TotalAmountLists,
} from "~/types";

const INITIAL_STORE_EXPENSE_INPUT = {
  paidAt: new Date(),
  amount: 0,
  category: 0,
  description: "",
};
const INITIAL_EXPENSE_VALIDATION_ERRORS = {
  amount: [],
  category: [],
  description: [],
};

const INITIAL_STORE_INCOME_INPUT = {
  receivedAt: new Date(),
  amount: 0,
  clientName: "",
};
const INITIAL_INCOME_VALIDATION_ERRORS = {
  amount: [],
  clientName: [],
};

export const useMonthlyBudgetCalender = () => {
  const [events, setEvents] = useState<TotalAmountLists>([]);
  const [inView, setInView] = useState(false);
  const [selectedDate, setSelectedDate] = useState("");

  const now = new Date();
  const [selectedMonth, setSelectedMonth] = useState<{
    beginning: Date;
    end: Date;
  }>({
    beginning: new Date(now.getFullYear(), now.getMonth(), 1),
    end: new Date(now.getFullYear(), now.getMonth() + 1, 0),
  });

  // NOTE: 支出登録関連
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
  const [expenseValidationErrors, setExpenseValidationErrors] = useState<StoreExpenseValidationError>(INITIAL_EXPENSE_VALIDATION_ERRORS);

  // NOTE: 収入登録関連
  const [storeIncomeInput, setStoreIncomeInput] = useState<StoreIncomeInput>(INITIAL_STORE_INCOME_INPUT);
  const updateStoreIncomeInput = useCallback((params: Partial<StoreIncomeInput>) => {
    setStoreIncomeInput((prev: StoreIncomeInput) => ({ ...prev, ...params }));
  }, []);
  const setStoreIncomeTextInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      updateStoreIncomeInput({ [e.target.name]: e.target.value });
    },
    [updateStoreIncomeInput],
  );
  const [incomeValidationErrors, setIncomeValidationErrors] = useState<StoreIncomeValidationError>(INITIAL_INCOME_VALIDATION_ERRORS);

  // NOTE: 月が変更された時の処理
  const handleDatesSet = async (arg: DatesSetArg) => {
    const selectedMonthBeginningDate = arg.view.currentStart;
    const currentEnd = arg.view.currentEnd;
    currentEnd.setDate(currentEnd.getDate() - 1);
    const selectedMonthEndDate = currentEnd;

    setSelectedMonth({
      beginning: selectedMonthBeginningDate,
      end: selectedMonthEndDate,
    });
  };

  // NOTE: 日が選択された時の処理
  const handleDateClick = async (arg: DateClickArg) => {
    // NOTE: 選択月以外の日付のクリックは無効にする
    if (arg.date.getMonth() !== selectedMonth.beginning.getMonth()) return;

    const date = getDateString(arg.date);
    setSelectedDate(date);
    setInView(true);
    setStoreExpenseInput((prev: StoreExpenseInput) => ({ ...prev, ...{ paidAt: arg.date } }));
    setStoreIncomeInput((prev: StoreIncomeInput) => ({ ...prev, ...{ receivedAt: arg.date } }));
  };

  const initExpenseValidationErrors = useCallback(() => {
    setExpenseValidationErrors(INITIAL_EXPENSE_VALIDATION_ERRORS);
  }, []);

  const onSuccessPostCreateExpense = useCallback(
    (result: StoreExpenseResponse) => {
      if (Object.keys(result.errors).length > 0) {
        setExpenseValidationErrors(result.errors);
        return;
      }

      window.alert("支出を登録しました");

      // NOTE: 選択月の支出に反映する
      const otherDateEvents = events.filter((event) => String(event.date) !== selectedDate || event.extendProps.type !== "expense");
      const currentDateEvent = events.find((event) => String(event.date) === selectedDate && event.extendProps.type === "expense");
      const newTotalAmount = currentDateEvent?.extendProps.totalAmount ?? 0;
      setEvents([
        ...otherDateEvents,
        {
          date: result.expense.paidAt,
          extendProps: {
            type: "expense",
            totalAmount: newTotalAmount + result.expense.amount,
          },
        },
      ]);
      setInView(false);
      setStoreExpenseInput(INITIAL_STORE_EXPENSE_INPUT);
      setStoreIncomeInput(INITIAL_STORE_INCOME_INPUT);
    },
    [events, selectedDate, setEvents, setInView, setStoreExpenseInput, setStoreIncomeInput],
  );

  const initIncomeValidationErrors = useCallback(() => {
    setIncomeValidationErrors(INITIAL_INCOME_VALIDATION_ERRORS);
  }, []);

  const onSuccessPostCreateIncome = useCallback(
    (result: StoreIncomeResponse) => {
      if (Object.keys(result.errors).length > 0) {
        setIncomeValidationErrors(result.errors);
        return;
      }

      window.alert("収入を登録しました");

      // NOTE: 選択月の収入に反映する
      const otherDateEvents = events.filter((event) => String(event.date) !== selectedDate || event.extendProps.type !== "income");
      const currentDateEvent = events.find((event) => String(event.date) === selectedDate && event.extendProps.type === "income");
      const newTotalAmount = currentDateEvent?.extendProps.totalAmount ?? 0;
      setEvents([
        ...otherDateEvents,
        {
          date: result.income.receivedAt,
          extendProps: {
            type: "income",
            totalAmount: newTotalAmount + result.income.amount,
          },
        },
      ]);
      setInView(false);
      setStoreExpenseInput(INITIAL_STORE_EXPENSE_INPUT);
      setStoreIncomeInput(INITIAL_STORE_INCOME_INPUT);
    },
    [events, selectedDate, setEvents, setInView, setStoreExpenseInput, setStoreIncomeInput],
  );

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
    selectedMonth,
    summary,
    handleDatesSet,
    handleDateClick,
    setEvents,
    events,

    dialog: {
      inView,
      setInView,
      selectedDate,
      store: {
        // NOTE: 支出登録関連
        initExpenseValidationErrors,
        onSuccessPostCreateExpense,
        storeExpenseInput,
        expenseValidationErrors,
        setStoreExpenseTextInput,
        setStoreExpenseSelectInput,

        // NOTE: 収入登録関連
        initIncomeValidationErrors,
        onSuccessPostCreateIncome,
        storeIncomeInput,
        incomeValidationErrors,
        setStoreIncomeTextInput,
      },
    },
  };
};

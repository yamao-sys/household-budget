import type { DatesSetArg } from "@fullcalendar/core/index.js";
import type { DateClickArg } from "@fullcalendar/interaction/index.js";
import { useCallback, useMemo, useState } from "react";
import { getTotalAmounts } from "~/apis/expenses.api";
import { getIncomes, getIncomeTotalAmounts, postCreateIncome } from "~/apis/incomes.api";
import { useAuthContext } from "~/contexts/useAuthContext";
import { getDateString } from "~/lib/date";
import type {
  Income,
  StoreExpenseInput,
  StoreExpenseResponse,
  StoreExpenseValidationError,
  StoreIncomeInput,
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
  const [selectedDateIncomes, setSelectedDateIncomes] = useState<Income[]>([]);
  const [currentMonthDate, setCurrentMonthDate] = useState<Date>(new Date());

  const { csrfToken } = useAuthContext();

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

    // TODO: ここをTanstack Query等を使用してキャッシュする
    const fetchedExpenseTotalAmounts = await getTotalAmounts(
      getDateString(selectedMonthBeginningDate),
      getDateString(selectedMonthEndDate),
      csrfToken,
    );
    const fetchedIncomeTotalAmounts = await getIncomeTotalAmounts(
      getDateString(selectedMonthBeginningDate),
      getDateString(selectedMonthEndDate),
      csrfToken,
    );

    setCurrentMonthDate(selectedMonthBeginningDate);
    setEvents([...(fetchedExpenseTotalAmounts ?? []), ...(fetchedIncomeTotalAmounts ?? [])]);
  };

  // NOTE: 日が選択された時の処理
  const handleDateClick = async (arg: DateClickArg) => {
    // NOTE: 選択月以外の日付のクリックは無効にする
    if (arg.date.getMonth() !== currentMonthDate.getMonth()) return;

    const date = getDateString(arg.date);
    setSelectedDateIncomes(await getIncomes(date, date, csrfToken));
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

  const handleCreateIncome = async () => {
    setIncomeValidationErrors(INITIAL_INCOME_VALIDATION_ERRORS);

    const { income, errors } = await postCreateIncome(storeIncomeInput, csrfToken);
    if (Object.keys(errors).length > 0) {
      setIncomeValidationErrors(errors);
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
        date: income.receivedAt,
        extendProps: {
          type: "income",
          totalAmount: newTotalAmount + income.amount,
        },
      },
    ]);
    setInView(false);
    setStoreExpenseInput(INITIAL_STORE_EXPENSE_INPUT);
    setStoreIncomeInput(INITIAL_STORE_INCOME_INPUT);
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
      selectedDateIncomes,
      store: {
        // 支出登録
        initExpenseValidationErrors,
        onSuccessPostCreateExpense,
        storeExpenseInput,
        expenseValidationErrors,
        setStoreExpenseTextInput,
        setStoreExpenseSelectInput,

        // 収入登録
        handleCreateIncome,
        storeIncomeInput,
        incomeValidationErrors,
        setStoreIncomeTextInput,
      },
    },
  };
};

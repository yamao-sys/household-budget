import { getRequestHeaders } from "../base/api";
import type {
  CategoryTotalAmountLists,
  GetExpensesCategoryTotalAmountsParams,
  GetExpensesParams,
  GetExpensesTotalAmountsParams,
  StoreExpenseInput,
  StoreExpenseResponse,
  TotalAmountLists,
} from "~/apis/model";
import {
  getExpenses as getExpensesApi,
  getExpensesTotalAmounts as getExpenseTotalAmountsApi,
  getExpensesCategoryTotalAmounts as getExpensesCategoryTotalAmountsApi,
  postExpenses,
} from "~/apis/expenses/expenses";

export const getExpenses = async (fromDate: string, toDate: string, csrfToken: string) => {
  const params: GetExpensesParams = {};
  if (!!fromDate) {
    params.fromDate = fromDate;
  }
  if (!!toDate) {
    params.toDate = toDate;
  }
  try {
    const res = await getExpensesApi(params, getRequestHeaders(csrfToken));

    return res.data.expenses;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
};

export async function getExpenseTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<TotalAmountLists[]> {
  const params: GetExpensesTotalAmountsParams = { fromDate, toDate };
  try {
    const res = await getExpenseTotalAmountsApi(params, getRequestHeaders(csrfToken));

    return res.data.totalAmounts;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

export async function getExpenseCategoryTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<CategoryTotalAmountLists[]> {
  const params: GetExpensesCategoryTotalAmountsParams = { fromDate, toDate };
  try {
    const res = await getExpensesCategoryTotalAmountsApi(params, getRequestHeaders(csrfToken));

    return res.data.totalAmounts;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

export async function postCreateExpense(input: StoreExpenseInput, csrfToken: string): Promise<StoreExpenseResponse> {
  try {
    const res = await postExpenses(input, getRequestHeaders(csrfToken));

    if (res.status === 500) {
      throw new Error(`Internal Server Error: ${res.data}`);
    }

    return res.data;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

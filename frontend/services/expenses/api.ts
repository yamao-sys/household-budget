import type { operations } from "~/types/apiSchema";
import { client, getRequestHeaders } from "../base/api";
import type { CategoryTotalAmountLists, Expense, StoreExpenseInput, StoreExpenseResponse, TotalAmountLists } from "~/types";

export async function getExpenses(fromDate: string, toDate: string, csrfToken: string): Promise<Expense[]> {
  const params: operations["get-expenses"]["parameters"] = { query: {} };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }
  const emptyExpenses: Expense[] = [];

  if (!fromDate && !toDate) return emptyExpenses;

  const { data } = await client.GET("/expenses", {
    ...getRequestHeaders(csrfToken),
    params,
  });

  return data?.expenses ?? emptyExpenses;
}

export async function getExpenseTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<TotalAmountLists> {
  const params: operations["get-expenses-total-amounts"]["parameters"] = { query: { fromDate: "", toDate: "" } };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/expenses/totalAmounts", {
    ...getRequestHeaders(csrfToken),
    params,
  });
  if (!data) {
    throw new Error();
  }

  return data.totalAmounts;
}

export async function getExpenseCategoryTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<CategoryTotalAmountLists> {
  const params: operations["get-expenses-category-total-amounts"]["parameters"] = { query: { fromDate: "", toDate: "" } };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/expenses/categoryTotalAmounts", {
    ...getRequestHeaders(csrfToken),
    params,
  });
  if (!data) {
    throw new Error();
  }

  return data.totalAmounts;
}

export async function postCreateExpense(input: StoreExpenseInput, csrfToken: string): Promise<StoreExpenseResponse> {
  const { data } = await client.POST("/expenses", {
    ...getRequestHeaders(csrfToken),
    body: input,
    bodySerializer() {
      const reqBody: { [key: string]: string | number } = {};
      for (const [key, value] of Object.entries(input)) {
        if (value instanceof Date) {
          reqBody[key] = value.toLocaleDateString("ja-JP", { year: "numeric", month: "2-digit", day: "2-digit" }).replaceAll("/", "-");
        } else if (["amount", "category"].includes(key)) {
          if (value) {
            reqBody[key] = Number(value);
          }
        } else {
          reqBody[key] = value;
        }
      }
      return JSON.stringify(reqBody);
    },
  });
  if (!data) {
    throw new Error();
  }

  return data;
}

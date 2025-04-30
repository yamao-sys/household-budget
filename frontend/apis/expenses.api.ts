import createClient from "openapi-fetch";
import type { operations, paths } from "./generated/apiSchema";
import { getRequestHeaders } from "./csrf.api";
import type { Expense, StoreExpenseInput } from "~/types";

const client = createClient<paths>({
  baseUrl: `${import.meta.env.VITE_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function getExpenses(fromDate: string, toDate: string) {
  const params: operations["get-expenses"]["parameters"] = { query: {} };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/expenses", {
    ...(await getRequestHeaders()),
    params,
  });

  const emptyExpenses: Expense[] = [];

  return data?.expenses ?? emptyExpenses;
}

export async function getTotalAmounts(fromDate: string, toDate: string) {
  const params: operations["get-expenses-total-amounts"]["parameters"] = { query: { fromDate: "", toDate: "" } };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/expenses/totalAmounts", {
    ...(await getRequestHeaders()),
    params,
  });
  if (!data) {
    throw new Error();
  }

  return data.totalAmounts;
}

export async function getCategoryTotalAmounts(fromDate: string, toDate: string) {
  const params: operations["get-expenses-category-total-amounts"]["parameters"] = { query: { fromDate: "", toDate: "" } };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/expenses/categoryTotalAmounts", {
    ...(await getRequestHeaders()),
    params,
  });
  if (!data) {
    throw new Error();
  }

  return data.totalAmounts;
}

export async function postCreateExpense(input: StoreExpenseInput) {
  const { data } = await client.POST("/expenses", {
    ...(await getRequestHeaders()),
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

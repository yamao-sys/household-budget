import createClient from "openapi-fetch";
import type { operations, paths } from "./generated/apiSchema";
import { getRequestHeaders } from "./csrf.api";
import type { Expense } from "~/types";

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

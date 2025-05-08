import createClient from "openapi-fetch";
import type { operations, paths } from "./generated/apiSchema";
import { getRequestHeaders } from "./csrf.api";
import type { Income, StoreIncomeInput } from "~/types";

const client = createClient<paths>({
  baseUrl: `${import.meta.env.VITE_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function getIncomes(fromDate: string, toDate: string, csrfToken: string) {
  const params: operations["get-incomes"]["parameters"] = { query: {} };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/incomes", {
    ...getRequestHeaders(csrfToken),
    params,
  });

  const emptyIncomes: Income[] = [];

  return data?.incomes ?? emptyIncomes;
}

export async function getIncomeTotalAmounts(fromDate: string, toDate: string, csrfToken: string) {
  const params: operations["get-incomes-total-amounts"]["parameters"] = { query: { fromDate: "", toDate: "" } };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/incomes/totalAmounts", {
    ...getRequestHeaders(csrfToken),
    params,
  });
  if (!data) {
    throw new Error();
  }

  return data.totalAmounts;
}

export async function getClientTotalAmounts(fromDate: string, toDate: string, csrfToken: string) {
  const params: operations["get-incomes-client-total-amounts"]["parameters"] = { query: { fromDate: "", toDate: "" } };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const { data } = await client.GET("/incomes/clientTotalAmounts", {
    ...getRequestHeaders(csrfToken),
    params,
  });
  if (!data) {
    throw new Error();
  }

  return data.totalAmounts;
}

export async function postCreateIncome(input: StoreIncomeInput, csrfToken: string) {
  const { data } = await client.POST("/incomes", {
    ...getRequestHeaders(csrfToken),
    body: input,
    bodySerializer() {
      const reqBody: { [key: string]: string | number } = {};
      for (const [key, value] of Object.entries(input)) {
        if (value instanceof Date) {
          reqBody[key] = value.toLocaleDateString("ja-JP", { year: "numeric", month: "2-digit", day: "2-digit" }).replaceAll("/", "-");
        } else if (["amount"].includes(key)) {
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

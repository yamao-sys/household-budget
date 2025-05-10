import type { operations } from "~/types/apiSchema";
import { client, getRequestHeaders } from "../base/api";
import type { ClientTotalAmountLists, Income, StoreIncomeInput, StoreIncomeResponse, TotalAmountLists } from "~/types";

export async function getIncomes(fromDate: string, toDate: string, csrfToken: string): Promise<Income[]> {
  const params: operations["get-incomes"]["parameters"] = { query: {} };
  if (!!fromDate) {
    params.query = { ...params.query, fromDate };
  }
  if (!!toDate) {
    params.query = { ...params.query, toDate };
  }

  const emptyIncomes: Income[] = [];

  if (!fromDate && !toDate) return emptyIncomes;

  const { data } = await client.GET("/incomes", {
    ...getRequestHeaders(csrfToken),
    params,
  });

  return data?.incomes ?? emptyIncomes;
}

export async function getIncomeTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<TotalAmountLists> {
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

export async function getIncomeClientTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<ClientTotalAmountLists> {
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

export async function postCreateIncome(input: StoreIncomeInput, csrfToken: string): Promise<StoreIncomeResponse> {
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

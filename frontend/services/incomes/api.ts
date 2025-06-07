import { getRequestHeaders } from "../base/api";
import { getIncomes as getIncomesApi, getIncomesClientTotalAmounts, getIncomesTotalAmounts, postIncomes } from "~/apis/incomes/incomes";
import type {
  ClientTotalAmountLists,
  GetIncomesClientTotalAmountsParams,
  GetIncomesParams,
  GetIncomesTotalAmountsParams,
  Income,
  StoreIncomeInput,
  StoreIncomeResponse,
  TotalAmountLists,
} from "~/apis/model";

export async function getIncomes(fromDate: string, toDate: string, csrfToken: string): Promise<Income[]> {
  const params: GetIncomesParams = {};
  if (!!fromDate) {
    params.fromDate = fromDate;
  }
  if (!!toDate) {
    params.toDate = toDate;
  }
  try {
    const res = await getIncomesApi(params, getRequestHeaders(csrfToken));

    return res.data.incomes;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

export async function getIncomeTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<TotalAmountLists[]> {
  const params: GetIncomesTotalAmountsParams = { fromDate, toDate };
  try {
    const res = await getIncomesTotalAmounts(params, getRequestHeaders(csrfToken));

    return res.data.totalAmounts;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

export async function getIncomeClientTotalAmounts(fromDate: string, toDate: string, csrfToken: string): Promise<ClientTotalAmountLists[]> {
  const params: GetIncomesClientTotalAmountsParams = { fromDate, toDate };
  try {
    const res = await getIncomesClientTotalAmounts(params, getRequestHeaders(csrfToken));

    return res.data.totalAmounts;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

export async function postCreateIncome(input: StoreIncomeInput, csrfToken: string): Promise<StoreIncomeResponse> {
  try {
    const res = await postIncomes(input, getRequestHeaders(csrfToken));

    if (res.status === 500) {
      throw new Error(`Internal Server Error: ${res.data}`);
    }

    return res.data;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

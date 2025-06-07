import type { components } from "~/types/apiSchema";

export type Income = components["schemas"]["Income"];

export type IncomeLists = components["schemas"]["IncomeLists"];

export type ClientTotalAmountLists = components["responses"]["ClientTotalAmountListsResponse"]["content"]["application/json"]["totalAmounts"];

export type StoreIncomeInput = components["requestBodies"]["StoreIncomeInput"]["content"]["application/json"];

export type StoreIncomeValidationError = components["responses"]["StoreIncomeResponse"]["content"]["application/json"]["errors"];

export type StoreIncomeResponse = components["responses"]["StoreIncomeResponse"]["content"]["application/json"];

import type { components } from "~/apis/generated/apiSchema";

export type UserSignUpInput = components["requestBodies"]["UserSignUpInput"]["content"]["application/json"];

export type UserSignUpValidationError = components["responses"]["UserSignUpResponse"]["content"]["application/json"]["errors"];

export type UserSignInInput = components["requestBodies"]["UserSignInInput"]["content"]["application/json"];

export type Expense = components["schemas"]["Expense"];

export type TotalAmountLists = components["responses"]["TotalAmountListsResponse"]["content"]["application/json"]["totalAmounts"];

export type CategoryTotalAmountLists = components["responses"]["CategoryTotalAmountListsResponse"]["content"]["application/json"]["totalAmounts"];

export type ExpenseLists = components["schemas"]["ExpenseLists"];

export type Income = components["schemas"]["Income"];

export type IncomeLists = components["schemas"]["IncomeLists"];

export type ClientTotalAmountLists = components["responses"]["ClientTotalAmountListsResponse"]["content"]["application/json"]["totalAmounts"];

export type StoreExpenseInput = components["requestBodies"]["StoreExpenseInput"]["content"]["application/json"];

export type StoreExpenseValidationError = components["responses"]["StoreExpenseResponse"]["content"]["application/json"]["errors"];

export type StoreIncomeInput = components["requestBodies"]["StoreIncomeInput"]["content"]["application/json"];

export type StoreIncomeValidationError = components["responses"]["StoreIncomeResponse"]["content"]["application/json"]["errors"];

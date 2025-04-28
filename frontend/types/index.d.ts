import type { components } from "~/apis/generated/apiSchema";

export type UserSignUpInput = components["requestBodies"]["UserSignUpInput"]["content"]["application/json"];

export type UserSignUpValidationError = components["responses"]["UserSignUpResponse"]["content"]["application/json"]["errors"];

export type UserSignInInput = components["requestBodies"]["UserSignInInput"]["content"]["application/json"];

export interface paths {
  "/csrf": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Csrf */
    get: operations["get-csrf"];
    put?: never;
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/users/validateSignUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** User Validate SignUp */
    post: operations["post-users-validate_sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/users/signUp": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** User SignUp */
    post: operations["post-users-sign_up"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/users/signIn": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    get?: never;
    put?: never;
    /** User SignIn */
    post: operations["post-users-sign_in"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/users/checkSignedIn": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** User CheckSignedIn */
    get: operations["get-users-check_signed_in"];
    put?: never;
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/expenses": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Expenses */
    get: operations["get-expenses"];
    put?: never;
    /** Post Expense */
    post: operations["post-expenses"];
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/expenses/totalAmounts": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Expenses TotalAmounts */
    get: operations["get-expenses-total-amounts"];
    put?: never;
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
  "/expenses/categoryTotalAmounts": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    /** Get Expenses Category TotalAmounts */
    get: operations["get-expenses-category-total-amounts"];
    put?: never;
    post?: never;
    delete?: never;
    options?: never;
    head?: never;
    patch?: never;
    trace?: never;
  };
}
export type webhooks = Record<string, never>;
export interface components {
  schemas: {
    /** UserSignUpValidationError */
    UserSignUpValidationError: {
      name?: string[];
      email?: string[];
      password?: string[];
    };
    /**
     * TotalAmountLists
     * @description Total Amount Lists
     */
    TotalAmountLists: {
      /** Format: date */
      date: Date;
      extendProps: {
        type: string;
        totalAmount: number;
      };
    };
    /**
     * CategoryTotalAmountLists
     * @description Total Amount Lists
     */
    CategoryTotalAmountLists: {
      category: number;
      totalAmount: number;
    };
    /**
     * ExpenseLists
     * @description Monthly Calender Expense
     */
    ExpenseLists: {
      expenses: components["schemas"]["Expense"][];
    };
    /** StoreExpenseValidationError */
    StoreExpenseValidationError: {
      paidAt?: string[];
      amount?: string[];
      category?: string[];
      description?: string[];
    };
    /**
     * Expense
     * @description Expense
     */
    Expense: {
      id: string;
      /** Format: date */
      paidAt: Date;
      amount: number;
      category: number;
      description: string;
    };
  };
  responses: {
    /** @description Csrf response */
    CsrfResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          csrfToken: string;
        };
      };
    };
    UserSignUpResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          code: number;
          errors: components["schemas"]["UserSignUpValidationError"];
        };
      };
    };
    /** @description User SignIn Response */
    UserSignInOkResponse: {
      headers: {
        "Set-Cookie"?: string;
        [name: string]: unknown;
      };
      content: {
        "application/json": Record<string, never>;
      };
    };
    /** @description User SignIn BadRequest Response */
    UserSignInBadRequestResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          errors: string[];
        };
      };
    };
    /** @description Fetch Expense Lists Response */
    FetchExpenseListsResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": components["schemas"]["ExpenseLists"];
      };
    };
    /** @description Total Amount Lists Response */
    TotalAmountListsResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          totalAmounts: components["schemas"]["TotalAmountLists"][];
        };
      };
    };
    /** @description Category Total Amount Lists Response */
    CategoryTotalAmountListsResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          totalAmounts: components["schemas"]["CategoryTotalAmountLists"][];
        };
      };
    };
    /** @description Store Expense Response */
    StoreExpenseResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          expense: components["schemas"]["Expense"];
          errors: components["schemas"]["StoreExpenseValidationError"];
        };
      };
    };
    /** @description Internal Server Error Response */
    InternalServerErrorResponse: {
      headers: {
        [name: string]: unknown;
      };
      content: {
        "application/json": {
          /** Format: int64 */
          code: number;
          message: string;
        };
      };
    };
  };
  parameters: never;
  requestBodies: {
    /** @description User SignUp Iuput */
    UserSignUpInput: {
      content: {
        "application/json": {
          name: string;
          email: string;
          password: string;
        };
      };
    };
    /** @description User SignIn  Input */
    UserSignInInput: {
      content: {
        "application/json": {
          email: string;
          password: string;
        };
      };
    };
    /** @description Store Expense Input */
    StoreExpenseInput: {
      content: {
        "application/json": {
          /** Format: date */
          paidAt: Date;
          amount: number;
          category: number;
          description: string;
        };
      };
    };
  };
  headers: never;
  pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
  "get-csrf": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["CsrfResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-users-validate_sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["UserSignUpInput"];
    responses: {
      200: components["responses"]["UserSignUpResponse"];
      400: components["responses"]["UserSignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-users-sign_up": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["UserSignUpInput"];
    responses: {
      200: components["responses"]["UserSignUpResponse"];
      400: components["responses"]["UserSignUpResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "post-users-sign_in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["UserSignInInput"];
    responses: {
      200: components["responses"]["UserSignInOkResponse"];
      400: components["responses"]["UserSignInBadRequestResponse"];
      500: components["responses"]["InternalServerErrorResponse"];
    };
  };
  "get-users-check_signed_in": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      /** @description is signed in */
      200: {
        headers: {
          [name: string]: unknown;
        };
        content: {
          "application/json": {
            isSignedIn: boolean;
          };
        };
      };
    };
  };
  "get-expenses": {
    parameters: {
      query?: {
        /** @description 取得対象の日付FROM */
        fromDate?: string;
        /** @description 取得対象の日付TO */
        toDate?: string;
      };
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["FetchExpenseListsResponse"];
    };
  };
  "post-expenses": {
    parameters: {
      query?: never;
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: components["requestBodies"]["StoreExpenseInput"];
    responses: {
      200: components["responses"]["StoreExpenseResponse"];
    };
  };
  "get-expenses-total-amounts": {
    parameters: {
      query: {
        /** @description 取得対象の日付FROM */
        fromDate: string;
        /** @description 取得対象の日付TO */
        toDate: string;
      };
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["TotalAmountListsResponse"];
    };
  };
  "get-expenses-category-total-amounts": {
    parameters: {
      query: {
        /** @description 取得対象の日付FROM */
        fromDate: string;
        /** @description 取得対象の日付TO */
        toDate: string;
      };
      header?: never;
      path?: never;
      cookie?: never;
    };
    requestBody?: never;
    responses: {
      200: components["responses"]["CategoryTotalAmountListsResponse"];
    };
  };
}

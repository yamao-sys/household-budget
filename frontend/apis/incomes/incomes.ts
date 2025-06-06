/**
 * Generated by orval v7.9.0 🍺
 * Do not edit manually.
 * Household Budget Service
 * household budget
 * OpenAPI spec version: 1.0
 */
import { useMutation, useQuery } from "@tanstack/react-query";
import type {
  DataTag,
  DefinedInitialDataOptions,
  DefinedUseQueryResult,
  MutationFunction,
  QueryClient,
  QueryFunction,
  QueryKey,
  UndefinedInitialDataOptions,
  UseMutationOptions,
  UseMutationResult,
  UseQueryOptions,
  UseQueryResult,
} from "@tanstack/react-query";

import type {
  ClientTotalAmountListsResponse,
  GetIncomesClientTotalAmountsParams,
  GetIncomesParams,
  GetIncomesTotalAmountsParams,
  IncomeLists,
  PostIncomes500,
  StoreIncomeInput,
  StoreIncomeResponse,
  TotalAmountListsResponse,
} from ".././model";

import { customFetch } from "../../custom-fetch";

type SecondParameter<T extends (...args: never) => unknown> = Parameters<T>[1];

/**
 * @summary Get Incomes
 */
export type getIncomesResponse200 = {
  data: IncomeLists;
  status: 200;
};

export type getIncomesResponseComposite = getIncomesResponse200;

export type getIncomesResponse = getIncomesResponseComposite & {
  headers: Headers;
};

export const getGetIncomesUrl = (params?: GetIncomesParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value !== undefined) {
      normalizedParams.append(key, value === null ? "null" : value.toString());
    }
  });

  const stringifiedParams = normalizedParams.toString();

  return stringifiedParams.length > 0 ? `/incomes?${stringifiedParams}` : `/incomes`;
};

export const getIncomes = async (params?: GetIncomesParams, options?: RequestInit): Promise<getIncomesResponse> => {
  return customFetch<getIncomesResponse>(getGetIncomesUrl(params), {
    ...options,
    method: "GET",
  });
};

export const getGetIncomesQueryKey = (params?: GetIncomesParams) => {
  return [`/incomes`, ...(params ? [params] : [])] as const;
};

export const getGetIncomesQueryOptions = <TData = Awaited<ReturnType<typeof getIncomes>>, TError = unknown>(
  params?: GetIncomesParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomes>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetIncomesQueryKey(params);

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getIncomes>>> = ({ signal }) => getIncomes(params, { signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseQueryOptions<Awaited<ReturnType<typeof getIncomes>>, TError, TData> & {
    queryKey: DataTag<QueryKey, TData, TError>;
  };
};

export type GetIncomesQueryResult = NonNullable<Awaited<ReturnType<typeof getIncomes>>>;
export type GetIncomesQueryError = unknown;

export function useGetIncomes<TData = Awaited<ReturnType<typeof getIncomes>>, TError = unknown>(
  params: undefined | GetIncomesParams,
  options: {
    query: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomes>>, TError, TData>> &
      Pick<DefinedInitialDataOptions<Awaited<ReturnType<typeof getIncomes>>, TError, Awaited<ReturnType<typeof getIncomes>>>, "initialData">;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): DefinedUseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
export function useGetIncomes<TData = Awaited<ReturnType<typeof getIncomes>>, TError = unknown>(
  params?: GetIncomesParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomes>>, TError, TData>> &
      Pick<UndefinedInitialDataOptions<Awaited<ReturnType<typeof getIncomes>>, TError, Awaited<ReturnType<typeof getIncomes>>>, "initialData">;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
export function useGetIncomes<TData = Awaited<ReturnType<typeof getIncomes>>, TError = unknown>(
  params?: GetIncomesParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomes>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
/**
 * @summary Get Incomes
 */

export function useGetIncomes<TData = Awaited<ReturnType<typeof getIncomes>>, TError = unknown>(
  params?: GetIncomesParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomes>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> } {
  const queryOptions = getGetIncomesQueryOptions(params, options);

  const query = useQuery(queryOptions, queryClient) as UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };

  query.queryKey = queryOptions.queryKey;

  return query;
}

/**
 * @summary Post Income
 */
export type postIncomesResponse200 = {
  data: StoreIncomeResponse;
  status: 200;
};

export type postIncomesResponse500 = {
  data: PostIncomes500;
  status: 500;
};

export type postIncomesResponseComposite = postIncomesResponse200 | postIncomesResponse500;

export type postIncomesResponse = postIncomesResponseComposite & {
  headers: Headers;
};

export const getPostIncomesUrl = () => {
  return `/incomes`;
};

export const postIncomes = async (storeIncomeInput: StoreIncomeInput, options?: RequestInit): Promise<postIncomesResponse> => {
  return customFetch<postIncomesResponse>(getPostIncomesUrl(), {
    ...options,
    method: "POST",
    headers: { "Content-Type": "application/json", ...options?.headers },
    body: JSON.stringify(storeIncomeInput),
  });
};

export const getPostIncomesMutationOptions = <TError = PostIncomes500, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof postIncomes>>, TError, { data: StoreIncomeInput }, TContext>;
  request?: SecondParameter<typeof customFetch>;
}): UseMutationOptions<Awaited<ReturnType<typeof postIncomes>>, TError, { data: StoreIncomeInput }, TContext> => {
  const mutationKey = ["postIncomes"];
  const { mutation: mutationOptions, request: requestOptions } = options
    ? options.mutation && "mutationKey" in options.mutation && options.mutation.mutationKey
      ? options
      : { ...options, mutation: { ...options.mutation, mutationKey } }
    : { mutation: { mutationKey }, request: undefined };

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof postIncomes>>, { data: StoreIncomeInput }> = (props) => {
    const { data } = props ?? {};

    return postIncomes(data, requestOptions);
  };

  return { mutationFn, ...mutationOptions };
};

export type PostIncomesMutationResult = NonNullable<Awaited<ReturnType<typeof postIncomes>>>;
export type PostIncomesMutationBody = StoreIncomeInput;
export type PostIncomesMutationError = PostIncomes500;

/**
 * @summary Post Income
 */
export const usePostIncomes = <TError = PostIncomes500, TContext = unknown>(
  options?: {
    mutation?: UseMutationOptions<Awaited<ReturnType<typeof postIncomes>>, TError, { data: StoreIncomeInput }, TContext>;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseMutationResult<Awaited<ReturnType<typeof postIncomes>>, TError, { data: StoreIncomeInput }, TContext> => {
  const mutationOptions = getPostIncomesMutationOptions(options);

  return useMutation(mutationOptions, queryClient);
};
/**
 * @summary Get Income Client TotalAmounts
 */
export type getIncomesClientTotalAmountsResponse200 = {
  data: ClientTotalAmountListsResponse;
  status: 200;
};

export type getIncomesClientTotalAmountsResponseComposite = getIncomesClientTotalAmountsResponse200;

export type getIncomesClientTotalAmountsResponse = getIncomesClientTotalAmountsResponseComposite & {
  headers: Headers;
};

export const getGetIncomesClientTotalAmountsUrl = (params: GetIncomesClientTotalAmountsParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value !== undefined) {
      normalizedParams.append(key, value === null ? "null" : value.toString());
    }
  });

  const stringifiedParams = normalizedParams.toString();

  return stringifiedParams.length > 0 ? `/incomes/clientTotalAmounts?${stringifiedParams}` : `/incomes/clientTotalAmounts`;
};

export const getIncomesClientTotalAmounts = async (
  params: GetIncomesClientTotalAmountsParams,
  options?: RequestInit,
): Promise<getIncomesClientTotalAmountsResponse> => {
  return customFetch<getIncomesClientTotalAmountsResponse>(getGetIncomesClientTotalAmountsUrl(params), {
    ...options,
    method: "GET",
  });
};

export const getGetIncomesClientTotalAmountsQueryKey = (params: GetIncomesClientTotalAmountsParams) => {
  return [`/incomes/clientTotalAmounts`, ...(params ? [params] : [])] as const;
};

export const getGetIncomesClientTotalAmountsQueryOptions = <TData = Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError = unknown>(
  params: GetIncomesClientTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetIncomesClientTotalAmountsQueryKey(params);

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>> = ({ signal }) =>
    getIncomesClientTotalAmounts(params, { signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseQueryOptions<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError, TData> & {
    queryKey: DataTag<QueryKey, TData, TError>;
  };
};

export type GetIncomesClientTotalAmountsQueryResult = NonNullable<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>>;
export type GetIncomesClientTotalAmountsQueryError = unknown;

export function useGetIncomesClientTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError = unknown>(
  params: GetIncomesClientTotalAmountsParams,
  options: {
    query: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError, TData>> &
      Pick<
        DefinedInitialDataOptions<
          Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>,
          TError,
          Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>
        >,
        "initialData"
      >;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): DefinedUseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
export function useGetIncomesClientTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError = unknown>(
  params: GetIncomesClientTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError, TData>> &
      Pick<
        UndefinedInitialDataOptions<
          Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>,
          TError,
          Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>
        >,
        "initialData"
      >;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
export function useGetIncomesClientTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError = unknown>(
  params: GetIncomesClientTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
/**
 * @summary Get Income Client TotalAmounts
 */

export function useGetIncomesClientTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError = unknown>(
  params: GetIncomesClientTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesClientTotalAmounts>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> } {
  const queryOptions = getGetIncomesClientTotalAmountsQueryOptions(params, options);

  const query = useQuery(queryOptions, queryClient) as UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };

  query.queryKey = queryOptions.queryKey;

  return query;
}

/**
 * @summary Get Income Total Amounts
 */
export type getIncomesTotalAmountsResponse200 = {
  data: TotalAmountListsResponse;
  status: 200;
};

export type getIncomesTotalAmountsResponseComposite = getIncomesTotalAmountsResponse200;

export type getIncomesTotalAmountsResponse = getIncomesTotalAmountsResponseComposite & {
  headers: Headers;
};

export const getGetIncomesTotalAmountsUrl = (params: GetIncomesTotalAmountsParams) => {
  const normalizedParams = new URLSearchParams();

  Object.entries(params || {}).forEach(([key, value]) => {
    if (value !== undefined) {
      normalizedParams.append(key, value === null ? "null" : value.toString());
    }
  });

  const stringifiedParams = normalizedParams.toString();

  return stringifiedParams.length > 0 ? `/incomes/totalAmounts?${stringifiedParams}` : `/incomes/totalAmounts`;
};

export const getIncomesTotalAmounts = async (
  params: GetIncomesTotalAmountsParams,
  options?: RequestInit,
): Promise<getIncomesTotalAmountsResponse> => {
  return customFetch<getIncomesTotalAmountsResponse>(getGetIncomesTotalAmountsUrl(params), {
    ...options,
    method: "GET",
  });
};

export const getGetIncomesTotalAmountsQueryKey = (params: GetIncomesTotalAmountsParams) => {
  return [`/incomes/totalAmounts`, ...(params ? [params] : [])] as const;
};

export const getGetIncomesTotalAmountsQueryOptions = <TData = Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError = unknown>(
  params: GetIncomesTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {};

  const queryKey = queryOptions?.queryKey ?? getGetIncomesTotalAmountsQueryKey(params);

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getIncomesTotalAmounts>>> = ({ signal }) =>
    getIncomesTotalAmounts(params, { signal, ...requestOptions });

  return { queryKey, queryFn, ...queryOptions } as UseQueryOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, TData> & {
    queryKey: DataTag<QueryKey, TData, TError>;
  };
};

export type GetIncomesTotalAmountsQueryResult = NonNullable<Awaited<ReturnType<typeof getIncomesTotalAmounts>>>;
export type GetIncomesTotalAmountsQueryError = unknown;

export function useGetIncomesTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError = unknown>(
  params: GetIncomesTotalAmountsParams,
  options: {
    query: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, TData>> &
      Pick<
        DefinedInitialDataOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, Awaited<ReturnType<typeof getIncomesTotalAmounts>>>,
        "initialData"
      >;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): DefinedUseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
export function useGetIncomesTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError = unknown>(
  params: GetIncomesTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, TData>> &
      Pick<
        UndefinedInitialDataOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, Awaited<ReturnType<typeof getIncomesTotalAmounts>>>,
        "initialData"
      >;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
export function useGetIncomesTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError = unknown>(
  params: GetIncomesTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };
/**
 * @summary Get Income Total Amounts
 */

export function useGetIncomesTotalAmounts<TData = Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError = unknown>(
  params: GetIncomesTotalAmountsParams,
  options?: {
    query?: Partial<UseQueryOptions<Awaited<ReturnType<typeof getIncomesTotalAmounts>>, TError, TData>>;
    request?: SecondParameter<typeof customFetch>;
  },
  queryClient?: QueryClient,
): UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> } {
  const queryOptions = getGetIncomesTotalAmountsQueryOptions(params, options);

  const query = useQuery(queryOptions, queryClient) as UseQueryResult<TData, TError> & { queryKey: DataTag<QueryKey, TData, TError> };

  query.queryKey = queryOptions.queryKey;

  return query;
}

import createClient from "openapi-fetch";
import type { paths } from "./generated/apiSchema";

const client = createClient<paths>({
  baseUrl: `${import.meta.env.VITE_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export async function getCsrfToken() {
  const { data, error } = await client.GET("/csrf");
  if (error?.code === 500 || data === undefined) {
    throw Error();
  }
  return data.csrfToken;
}

export const getRequestHeaders = (csrfToken: string) => {
  return {
    headers: {
      "X-CSRF-Token": csrfToken,
    },
  };
};

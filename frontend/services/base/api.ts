import createClient from "openapi-fetch";
import type { paths } from "~/apis/generated/apiSchema";

export const client = createClient<paths>({
  baseUrl: `${import.meta.env.VITE_API_ENDPOINT_URI}/`,
  credentials: "include",
});

export const getRequestHeaders = (csrfToken: string) => {
  return {
    headers: {
      "X-CSRF-Token": csrfToken,
    },
  };
};

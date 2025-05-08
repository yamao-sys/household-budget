import createClient from "openapi-fetch";
import type { paths } from "./generated/apiSchema";
import Cookies from "js-cookie";

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

export const getRequestHeaders = async () => {
  console.log(`Cookies.get("_csrf"): ${Cookies.get("_csrf")}`);
  return {
    headers: {
      "X-CSRF-Token": Cookies.get("_csrf"),
    },
  };
};

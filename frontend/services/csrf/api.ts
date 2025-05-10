import { client } from "../base/api";

export async function getCsrfToken() {
  const { data, error } = await client.GET("/csrf");
  if (error?.code === 500 || data === undefined) {
    throw Error();
  }
  return data.csrfToken;
}

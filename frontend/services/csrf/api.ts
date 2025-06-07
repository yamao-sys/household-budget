import { getCsrf } from "~/apis/csrf/csrf";

export async function getCsrfToken() {
  try {
    const res = await getCsrf();

    if (res.status === 500) {
      throw new Error(`Internal Server Error: ${res.data}`);
    }

    return res.data.csrfToken;
  } catch (error) {
    throw new Error(`Unexpected error: ${error}`);
  }
}

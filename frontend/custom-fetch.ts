// NOTE: Supports cases where `content-type` is other than `json`
const getBody = <T>(c: Response | Request): Promise<T> => {
  const contentType = c.headers.get("content-type");

  if (contentType && contentType.includes("application/json")) {
    return c.json();
  }

  if (contentType && contentType.includes("application/pdf")) {
    return c.blob() as Promise<T>;
  }

  return c.text() as Promise<T>;
};

// NOTE: Update just base url
const getUrl = (contextUrl: string): string => {
  const baseUrl = process.env.VITE_API_ENDPOINT_URI;

  if (!baseUrl) {
    throw new Error("baseUrl is not defined");
  }

  // NOTE: 相対パスを処理できるようにする
  const requestUrl = new URL(contextUrl, baseUrl);
  return requestUrl.toString();
};

// NOTE: headers
const getHeaders = (headers?: HeadersInit): HeadersInit => {
  return {
    ...headers,
  };
};

export const customFetch = async <T>(url: string, options: RequestInit): Promise<T> => {
  const requestUrl = getUrl(url);
  const requestHeaders = getHeaders(options.headers);

  const requestInit: RequestInit = {
    ...options,
    headers: requestHeaders,
  };

  const response = await fetch(requestUrl, requestInit);
  const data = await getBody<T>(response);

  return { status: response.status, data, headers: response.headers } as T;
};

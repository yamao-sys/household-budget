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

const formatToJstDate = (isoString: string): string => {
  const date = new Date(isoString);
  const jst = new Date(date.getTime() + 9 * 60 * 60 * 1000); // UTC+9

  const year = jst.getFullYear();
  const month = String(jst.getMonth() + 1).padStart(2, "0");
  const day = String(jst.getDate()).padStart(2, "0");

  return `${year}-${month}-${day}`;
};

const convertDatesToDateOnly = (input: any): any => {
  if (typeof input === "string" && /^\d{4}-\d{2}-\d{2}T/.test(input)) {
    return formatToJstDate(input);
  }

  if (Array.isArray(input)) {
    return input.map(convertDatesToDateOnly);
  }

  if (input && typeof input === "object") {
    return Object.entries(input).reduce((acc, [key, value]) => {
      acc[key] = convertDatesToDateOnly(value);
      return acc;
    }, {} as any);
  }

  return input;
};

export const customFetch = async <T>(url: string, options: RequestInit): Promise<T> => {
  const requestUrl = getUrl(url);
  const requestHeaders = getHeaders(options.headers);

  const normalizedHeaders = new Headers(requestHeaders);

  let body = options.body;

  if (typeof body === "string" && normalizedHeaders.get("Content-Type")?.includes("application/json")) {
    try {
      const parsed = JSON.parse(body);
      const converted = convertDatesToDateOnly(parsed);
      body = JSON.stringify(converted);
      console.log("Converted body:", body);
    } catch (err) {
      console.warn("Failed to parse or convert body:", err);
    }
  }

  const requestInit: RequestInit = {
    ...options,
    headers: requestHeaders,
    credentials: "include",
    body,
  };

  const response = await fetch(requestUrl, requestInit);
  const data = await getBody<T>(response);

  return { status: response.status, data, headers: response.headers } as T;
};

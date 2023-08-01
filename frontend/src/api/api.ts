const BASE_URL = "http://localhost:8080"; // Replace with your server URL

async function fetchJSON<T>(
  url: string,
  method: string,
  data?: LogPayload | FormData
): Promise<T> {
  let body = null;
  if (data) {
    body = url.includes("/upload")
      ? (data as FormData)
      : `${(data as LogPayload).timestamp} ${(data as LogPayload).message}`;
  }

  const requestOptions: RequestInit = {
    method,
    body,
  };

  const response = await fetch(url, requestOptions);
  const responseData = await response.json();

  if (!response.ok) {
    throw new Error(responseData.error || "Something went wrong");
  }

  return responseData;
}

export interface LogPayload {
  message: string;
  timestamp: string;
}

export async function logPost(data: LogPayload): Promise<void> {
  return fetchJSON<void>(`${BASE_URL}/api/log`, "POST", data);
}

export async function logGet(): Promise<void> {
  return fetchJSON<void>(`${BASE_URL}/api/log`, "GET");
}

export async function uploadFile(file: File): Promise<void> {
  const formData = new FormData();
  formData.append("file", file);

  return fetchJSON<void>(`${BASE_URL}/api/upload`, "POST", formData);
}

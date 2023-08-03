const BASE_URL = "http://localhost:8484"; // Replace with your server URL

async function fetchJSON<T>(
  url: string,
  method: string,
  data?: string | FormData
): Promise<T> {
  let body = null;
  if (data) {
    body = url.includes("/upload") ? (data as FormData) : (data as string);
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

export interface LogResponse {
  ack: string;
  data: string;
}

export interface NilResponse {
  ack: string;
  data: null;
}

export async function logPost(data: string): Promise<NilResponse> {
  return fetchJSON<NilResponse>(`${BASE_URL}/api/log`, "POST", data);
}

export async function logGet(): Promise<LogResponse> {
  return fetchJSON<LogResponse>(`${BASE_URL}/api/log`, "GET");
}

export async function uploadFile(file: File): Promise<NilResponse> {
  const formData = new FormData();
  formData.append("file", file);

  return fetchJSON<NilResponse>(`${BASE_URL}/api/upload`, "POST", formData);
}

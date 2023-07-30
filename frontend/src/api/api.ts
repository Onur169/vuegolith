// apiService.ts
const BASE_URL = "http://localhost:8080"; // Replace with your server URL

async function fetchJSON<T>(
  url: string,
  method: string,
  data?: any
): Promise<T> {
  const requestOptions: RequestInit = {
    method,
    headers: {
      "Content-Type": "application/json",
    },
    body: data ? JSON.stringify(data) : undefined,
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

export async function log(data: LogPayload): Promise<void> {
  return fetchJSON<void>(`${BASE_URL}/api/log`, "POST", data);
}

export async function uploadFile(formData: FormData): Promise<void> {
  return fetchJSON<void>(`${BASE_URL}/api/upload`, "POST", formData);
}

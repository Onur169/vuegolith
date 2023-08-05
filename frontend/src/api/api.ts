const useSecureLocalDomain = import.meta.env.VITE_USE_SECURE_LOCAL_DOMAIN === 'true';
const BASE_URL = replacePortInURL(
  useSecureLocalDomain ? `https://vuegolith.local/` : `http://${window.location.host}`,
  useSecureLocalDomain ? '' : '8484',
);

function replacePortInURL(url: string, newPort: string): string {
  const u = new URL(url);
  u.port = newPort;
  return u.toString();
}

async function fetchJSON<T>(url: string, method: string, data?: string | FormData): Promise<T> {
  let body = null;
  if (data) {
    body = url.includes('/upload') ? (data as FormData) : (data as string);
  }

  const requestOptions: RequestInit = {
    method,
    body,
  };

  const response = await fetch(url, requestOptions);
  const responseData = await response.json();

  if (!response.ok) {
    throw new Error(responseData.error || 'Something went wrong');
  }

  return responseData;
}

export interface LogPayload {
  message: string;
  timestamp: string;
}

export interface DataStringResponse {
  ack: string;
  data: string;
}

export interface DataStringArrResponse {
  ack: string;
  data: string[];
}

export interface LogResponse extends DataStringResponse {}
export interface UploadResponse extends DataStringArrResponse {}

export interface NilResponse {
  ack: string;
  data: null;
}

export async function logPost(data: string): Promise<NilResponse> {
  return fetchJSON<NilResponse>(`${BASE_URL}api/log`, 'POST', data);
}

export async function logGet(): Promise<LogResponse> {
  return fetchJSON<LogResponse>(`${BASE_URL}api/log`, 'GET');
}

export async function uploadFile(file: File): Promise<NilResponse> {
  const formData = new FormData();
  formData.append('file', file);

  return fetchJSON<NilResponse>(`${BASE_URL}api/uploads`, 'POST', formData);
}

export async function uploadsGet(): Promise<UploadResponse> {
  return fetchJSON<UploadResponse>(`${BASE_URL}api/uploads`, 'GET');
}

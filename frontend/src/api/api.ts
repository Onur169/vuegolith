const useSecureLocalDomain = import.meta.env.VITE_USE_SECURE_LOCAL_DOMAIN === 'true';
export const baseUrl = replacePortInURL(
  useSecureLocalDomain ? `https://vuegolith.local/` : `http://${window.location.host}`,
  useSecureLocalDomain ? '' : '8484',
);

// todo: fetchJSON refactoren und in kleinere Module aufteilen
// In der App.vue müssen die promises auf ack === "success" checken für den erfolg

export function replacePortInURL(url: string, newPort: string): string {
  const u = new URL(url);
  u.port = newPort;
  return u.toString();
}

export interface UploadsPayload {
  file: string;
}

export async function fetchJSON<T>(
  url: string,
  method: string,
  data?: string | FormData | UploadsPayload,
): Promise<T> {
  let body = null;
  if (data && url.includes('/upload') && method === 'POST') {
    body = data as FormData;
  } else if (data && url.includes('/uploads') && method === 'DELETE') {
    body = JSON.stringify(data);
  } else {
    body = data as string;
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

export function isSecureLocalDomainActive() {
  return useSecureLocalDomain;
}

export interface LogPayload {
  message: string;
  timestamp: string;
}

export interface DataStringResponse {
  ack: string;
  data: string;
}

export interface DataArrResponse<T> {
  ack: string;
  data: T[];
}

export interface UploadFile {
  name: string;
  size: number;
  isDir: boolean;
}

export interface LogResponse extends DataStringResponse {}
export interface UploadResponse extends DataArrResponse<UploadFile> {}
export interface UploadDeleteResponse extends DataStringResponse {}

export interface NilResponse {
  ack: string;
  data: null;
}

export async function logPost(data: string): Promise<NilResponse> {
  return fetchJSON<NilResponse>(`${baseUrl}api/log`, 'POST', data);
}

export async function logGet(): Promise<LogResponse> {
  return fetchJSON<LogResponse>(`${baseUrl}api/log`, 'GET');
}

export async function uploadFile(file: File): Promise<NilResponse> {
  const formData = new FormData();
  formData.append('file', file);

  return fetchJSON<NilResponse>(`${baseUrl}api/uploads`, 'POST', formData);
}

export async function uploadsGet(): Promise<UploadResponse> {
  return fetchJSON<UploadResponse>(`${baseUrl}api/uploads`, 'GET');
}

export async function uploadsDelete(payload: UploadsPayload): Promise<UploadDeleteResponse> {
  return fetchJSON<UploadDeleteResponse>(`${baseUrl}api/uploads`, 'DELETE', payload);
}

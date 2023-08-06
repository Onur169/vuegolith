const useSecureLocalDomain = import.meta.env.VITE_USE_SECURE_LOCAL_DOMAIN === 'true';
export const baseUrl = replacePortInURL(
  useSecureLocalDomain ? `https://vuegolith.local/` : `http://${window.location.host}`,
  useSecureLocalDomain ? '' : '8484',
);

// todo: request refactoren und in kleinere Module aufteilen
// In der App.vue müssen die promises auf ack === "success" checken für den erfolg

export function replacePortInURL(url: string, newPort: string): string {
  const u = new URL(url);
  u.port = newPort;
  return u.toString();
}

export async function request<T, P extends PayloadConfig | null>(
  url: string,
  method: HttpMethod,
  payload?: P,
): Promise<T> {
  let body;

  if (payload?.treatAsBodyInit) {
    const { data } = payload;
    body = data as BodyInit;
  }

  if (payload?.convertToJsonString) {
    const { data } = payload;
    body = JSON.stringify(data);
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

export enum HttpMethod {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
  PATCH = 'PATCH',
  HEAD = 'HEAD',
  OPTIONS = 'OPTIONS',
  CONNECT = 'CONNECT',
  TRACE = 'TRACE',
}

export interface NilResponse {
  ack: string;
  data: null;
}

export interface DataStringResponse {
  ack: string;
  data: string;
}

export interface DataArrResponse<T> {
  ack: string;
  data: T[];
}

export interface PayloadConfig {
  convertToJsonString?: boolean;
  treatAsBodyInit?: boolean;
  data: unknown;
}

export interface UploadsPayload extends PayloadConfig {
  data: { file: string };
}

export interface LogPayload extends PayloadConfig {
  data: string;
}

export interface FormDataPayload extends PayloadConfig {
  data: FormData;
}

export interface UploadFile {
  name: string;
  size: number;
  isDir: boolean;
}

export interface LogResponse extends DataStringResponse {}
export interface UploadResponse extends DataArrResponse<UploadFile> {}
export interface UploadDeleteResponse extends DataStringResponse {}

export async function logPost(content: string): Promise<NilResponse> {
  const payloadData: LogPayload = {
    data: content,
    convertToJsonString: false,
    treatAsBodyInit: true,
  };
  return request<NilResponse, LogPayload>(`${baseUrl}api/log`, HttpMethod.POST, payloadData);
}

export async function logGet(): Promise<LogResponse> {
  return request<LogResponse, null>(`${baseUrl}api/log`, HttpMethod.GET);
}

export async function uploadFile(file: File): Promise<NilResponse> {
  const formData = new FormData();
  formData.append('file', file);

  const payload: FormDataPayload = {
    data: formData,
    treatAsBodyInit: true,
  };

  return request<NilResponse, FormDataPayload>(`${baseUrl}api/uploads`, HttpMethod.POST, payload);
}

export async function uploadsGet(): Promise<UploadResponse> {
  return request<UploadResponse, null>(`${baseUrl}api/uploads`, HttpMethod.GET);
}

export async function uploadsDelete(fileName: string): Promise<UploadDeleteResponse> {
  const payloadData: UploadsPayload = {
    data: { file: fileName },
    convertToJsonString: true,
  };

  return request<UploadDeleteResponse, UploadsPayload>(
    `${baseUrl}api/uploads`,
    HttpMethod.DELETE,
    payloadData,
  );
}

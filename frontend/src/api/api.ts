const useSecureLocalDomain = import.meta.env.VITE_USE_SECURE_LOCAL_DOMAIN === 'true';

export const baseUrl = replacePortInURL(
  useSecureLocalDomain ? `https://vuegolith.local/` : `http://${window.location.host}`,
  useSecureLocalDomain ? '' : '8484',
);

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

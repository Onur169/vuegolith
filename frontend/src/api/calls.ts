import {
  DataArrResponse,
  DataStringResponse,
  HttpMethod,
  NilResponse,
  PayloadConfig,
  apiBaseUrl,
  request,
} from './api';

// Payloads
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
}

// Response Interfaces
export interface LogResponse extends DataStringResponse {}
export interface UploadResponse extends DataArrResponse<UploadFile> {}
export interface UploadDeleteResponse extends DataStringResponse {}

// Call Promises
export async function logPost(content: string): Promise<NilResponse> {
  const payloadData: LogPayload = {
    data: content,
    treatAsBodyInit: true,
  };
  return request<NilResponse, LogPayload>(`api/log`, HttpMethod.POST, payloadData);
}

export async function logGet(): Promise<LogResponse> {
  return request<LogResponse, null>(`api/log`, HttpMethod.GET);
}

export async function uploadFile(file: File): Promise<NilResponse> {
  const formData = new FormData();
  formData.append('file', file);

  const payload: FormDataPayload = {
    data: formData,
    treatAsBodyInit: true,
  };

  return request<NilResponse, FormDataPayload>(`api/uploads`, HttpMethod.POST, payload);
}

export async function uploadFileWithProgress(
  file: File,
  onProgress: (percentage: number) => void,
): Promise<NilResponse> {
  const formData = new FormData();
  formData.append('file', file);

  return new Promise<NilResponse>((resolve, reject) => {
    const xhr = new XMLHttpRequest();

    xhr.upload.addEventListener('progress', event => {
      if (event.lengthComputable) {
        const percentage = (event.loaded / event.total) * 100;
        onProgress(percentage);
      }
    });

    xhr.onreadystatechange = () => {
      if (xhr.readyState === XMLHttpRequest.DONE) {
        if (xhr.status === 200) {
          resolve(JSON.parse(xhr.responseText));
        } else {
          reject(xhr.statusText || 'Something went wrong');
        }
      }
    };

    xhr.open('POST', `${apiBaseUrl}api/uploads`, true);
    xhr.send(formData);
  });
}

export async function uploadsGet(): Promise<UploadResponse> {
  return request<UploadResponse, null>(`api/uploads`, HttpMethod.GET);
}

export async function uploadsDelete(fileName: string): Promise<UploadDeleteResponse> {
  const payloadData: UploadsPayload = {
    data: { file: fileName },
    convertToJsonString: true,
  };

  return request<UploadDeleteResponse, UploadsPayload>(
    `api/uploads`,
    HttpMethod.DELETE,
    payloadData,
  );
}

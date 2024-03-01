import { FetchError } from '@/utils';

const BASE_URL = 'http://localhost/api/';

export const sendToApi = async (uri: string, config: RequestInit): Promise<Response> => {
  const url = new URL(uri, BASE_URL);
  const response = await fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      ...config?.headers,
    },
    ...config
  });

  if (!response.ok) {
    const err = new FetchError(`HTTP Status code: ${response.status}`, response, response.status);
    throw err;
  }

  return response
}

export const sendToApiExpectBody = async <T>(uri: string, config: RequestInit): Promise<T> => {
  const url = new URL(uri, BASE_URL);
  const response = await fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      ...config?.headers,
    },
    ...config
  });

  const responseBody = await response.json();
  if (response.ok) {
    return responseBody as T
  }
  const error = new Error(responseBody.message);
  return Promise.reject(error);
}
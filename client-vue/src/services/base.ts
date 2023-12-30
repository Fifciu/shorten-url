const BASE_URL = 'http://localhost/api/';

export const sendToApi = async <T>(uri: string, config: RequestInit, hasResponseBody: boolean = true) => {
  const url = new URL(uri, BASE_URL);
  const response = await fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      ...config?.headers,
    },
    ...config
  });

  if (!hasResponseBody) {
    return response
  }

  const responseBody = await response.json();
  if (response.ok) {
    return responseBody as T
  }
  const error = new Error(responseBody.message);
  return Promise.reject(error);
}
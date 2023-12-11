const BASE_URL = 'http://localhost:2222';

export const sendToApi = async <T>(uri: string, config: RequestInit) => {
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
import { BASE_URL } from '@app/config';
import type { APIErrorResp } from '@app/types/api';
import axios, { AxiosError } from 'axios';

export type ErrorObject = {
  message: string;
  statusCode: number;
};

const genericErrorMsg = 'Something went wrong';

const axiosInstance = axios.create({
  baseURL: `${BASE_URL}/api`,
  headers: {
    'Content-Type': 'application/json',
    'x-auth-token': localStorage.getItem('accessToken') || ''
  }
});

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error: AxiosError<APIErrorResp>) => {
    const m = error?.response?.data?.message || genericErrorMsg;
    const err = {
      message: m,
      statusCode: error.response?.status
    };
    return Promise.reject(err as ErrorObject);
  }
);

export const setSession = (token: string) => {
  localStorage.setItem('accessToken', token);
  axiosInstance.defaults.headers['x-auth-token'] = token;
};

export const clearSession = () => {
  localStorage.removeItem('accessToken');
  axiosInstance.defaults.headers['x-auth-token'] = '';
};

export default axiosInstance;

export const errorHandler = (
  error: unknown,
  onError: (errorObj: ErrorObject) => void
): void => {
  if (error instanceof Error) {
    return onError({
      message: error.message || genericErrorMsg,
      statusCode: 500
    });
  }

  const apiErr = error as ErrorObject;

  if (apiErr?.message) {
    return onError(apiErr);
  }

  if (typeof error === 'string') {
    return onError({ message: error || genericErrorMsg, statusCode: 500 });
  }

  return onError({ message: genericErrorMsg, statusCode: 500 });
};

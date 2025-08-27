import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL as string,
});

let isRefreshing = false;
let failedQueue: {
  resolve: (value?: unknown) => void;
  reject: (reason?: any) => void;
}[] = [];

const processQueue = (error: unknown, token: string | null = null) => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });
  failedQueue = [];
};

api.interceptors.request.use((config) => {
  const token = localStorage.getItem("accessToken");
  if (token) {
    config.headers = config.headers || {};
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

api.interceptors.response.use(
  (response: AxiosResponse) => response,
  async (error: AxiosError) => {
    const originalRequest = error.config as AxiosRequestConfig & {
      _retry?: boolean;
    };

    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then((token) => {
            if (originalRequest.headers) {
              originalRequest.headers.Authorization = `Bearer ${token}`;
            }
            return api(originalRequest);
          })
          .catch(Promise.reject);
      }

      originalRequest._retry = true;
      isRefreshing = true;

      try {
        const refreshToken = localStorage.getItem("refreshToken");
        const { data } = await axios.post(
          `${import.meta.env.VITE_API_BASE_URL as string}/auth/refresh`,
          { refreshToken },
        );
        const accessToken = (data as any).accessToken;
        localStorage.setItem("accessToken", accessToken);
        processQueue(null, accessToken);
        if (originalRequest.headers) {
          originalRequest.headers.Authorization = `Bearer ${accessToken}`;
        }
        return api(originalRequest);
      } catch (err) {
        processQueue(err, null);
        return Promise.reject(err);
      } finally {
        isRefreshing = false;
      }
    }

    return Promise.reject(error);
  },
);

export const get = async <T = unknown>(
  url: string,
  config?: AxiosRequestConfig,
): Promise<T> => {
  try {
    const { data } = await api.get<T>(url, config);
    return data;
  } catch (err) {
    const axiosErr = err as AxiosError;
    throw axiosErr.response?.data || axiosErr;
  }
};

export const post = async <T = unknown>(
  url: string,
  body?: unknown,
  config?: AxiosRequestConfig,
): Promise<T> => {
  try {
    const { data } = await api.post<T>(url, body, config);
    return data;
  } catch (err) {
    const axiosErr = err as AxiosError;
    throw axiosErr.response?.data || axiosErr;
  }
};

export default api;

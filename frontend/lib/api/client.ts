import axios, { AxiosError, AxiosInstance, InternalAxiosRequestConfig } from 'axios'
import { ApiError } from '@/types/api'

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:30000/api/v1'

class ApiClient {
  private client: AxiosInstance

  constructor() {
    this.client = axios.create({
      baseURL: API_BASE_URL,
      headers: {
        'Content-Type': 'application/json',
      },
      withCredentials: true, // For cookies
    })

    this.setupInterceptors()
  }

  private setupInterceptors() {
    // Request interceptor - Add auth token
    this.client.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        // Get token from localStorage
        const authStorage = localStorage.getItem('auth-storage')
        if (authStorage) {
          const { state } = JSON.parse(authStorage)
          const token = state?.tokens?.access_token

          if (token && config.headers) {
            config.headers.Authorization = `Bearer ${token}`
          }
        }
        return config
      },
      (error) => Promise.reject(error)
    )

    // Response interceptor - Handle token refresh
    this.client.interceptors.response.use(
      (response) => response,
      async (error: AxiosError<ApiError>) => {
        const originalRequest = error.config as InternalAxiosRequestConfig & { _retry?: boolean }

        // If error is 401 and we haven't retried yet
        if (error.response?.status === 401 && !originalRequest._retry) {
          originalRequest._retry = true

          try {
            const authStorage = localStorage.getItem('auth-storage')
            if (authStorage) {
              const { state } = JSON.parse(authStorage)
              const refreshToken = state?.tokens?.refresh_token

              if (refreshToken) {
                // Try to refresh token
                const response = await this.client.post('/auth/refresh', {
                  refresh_token: refreshToken,
                })

                const { access_token } = response.data

                // Update token in storage
                const updatedState = {
                  ...state,
                  tokens: {
                    ...state.tokens,
                    access_token,
                  },
                }
                localStorage.setItem(
                  'auth-storage',
                  JSON.stringify({ state: updatedState, version: 0 })
                )

                // Retry original request with new token
                if (originalRequest.headers) {
                  originalRequest.headers.Authorization = `Bearer ${access_token}`
                }
                return this.client(originalRequest)
              }
            }
          } catch (refreshError) {
            // Refresh failed, clear auth and redirect to login
            localStorage.removeItem('auth-storage')
            window.location.href = '/login'
            return Promise.reject(refreshError)
          }
        }

        return Promise.reject(error)
      }
    )
  }

  // HTTP methods
  async get<T>(url: string, config = {}) {
    const response = await this.client.get<T>(url, config)
    return response.data
  }

  async post<T>(url: string, data?: unknown, config = {}) {
    const response = await this.client.post<T>(url, data, config)
    return response.data
  }

  async put<T>(url: string, data?: unknown, config = {}) {
    const response = await this.client.put<T>(url, data, config)
    return response.data
  }

  async delete<T>(url: string, config = {}) {
    const response = await this.client.delete<T>(url, config)
    return response.data
  }

  async patch<T>(url: string, data?: unknown, config = {}) {
    const response = await this.client.patch<T>(url, data, config)
    return response.data
  }
}

export const apiClient = new ApiClient()

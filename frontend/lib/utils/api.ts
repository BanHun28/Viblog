/**
 * API utilities for Viblog frontend
 * Request/response helpers and API configuration
 */

import { getLocalStorage, StorageKeys } from './storage';
import { buildQueryString as buildQuery } from './url';

/**
 * API configuration
 */
export const API_CONFIG = {
  BASE_URL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:30000',
  API_VERSION: 'v1',
  TIMEOUT: 30000, // 30 seconds
  RETRY_ATTEMPTS: 3,
  RETRY_DELAY: 1000 // 1 second
} as const;

/**
 * Build full API URL
 */
export function buildApiUrl(endpoint: string): string {
  const cleanEndpoint = endpoint.startsWith('/') ? endpoint.slice(1) : endpoint;
  return `${API_CONFIG.BASE_URL}/api/${API_CONFIG.API_VERSION}/${cleanEndpoint}`;
}

/**
 * Get authorization header
 */
export function getAuthHeader(): Record<string, string> {
  const token = getLocalStorage<string>(StorageKeys.AUTH_TOKEN);

  if (token) {
    return {
      'Authorization': `Bearer ${token}`
    };
  }

  return {};
}

/**
 * Build request headers
 */
export function buildHeaders(customHeaders?: Record<string, string>): HeadersInit {
  return {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
    ...getAuthHeader(),
    ...customHeaders
  };
}

// Re-export buildQueryString from url.ts for backward compatibility
export { buildQueryString } from './url';

/**
 * Sleep utility for retry logic
 */
export function sleep(ms: number): Promise<void> {
  return new Promise(resolve => setTimeout(resolve, ms));
}

/**
 * Retry wrapper for API requests
 */
export async function retryRequest<T>(
  fn: () => Promise<T>,
  attempts: number = API_CONFIG.RETRY_ATTEMPTS,
  delay: number = API_CONFIG.RETRY_DELAY
): Promise<T> {
  try {
    return await fn();
  } catch (error: any) {
    if (attempts <= 1) {
      throw error;
    }

    // Don't retry on client errors (4xx)
    if (error.response?.status >= 400 && error.response?.status < 500) {
      throw error;
    }

    await sleep(delay);
    return retryRequest(fn, attempts - 1, delay * 2); // Exponential backoff
  }
}

/**
 * Check if response is JSON
 */
export function isJsonResponse(response: Response): boolean {
  const contentType = response.headers.get('content-type');
  return contentType?.includes('application/json') || false;
}

/**
 * Parse response body
 */
export async function parseResponseBody<T>(response: Response): Promise<T> {
  if (isJsonResponse(response)) {
    return response.json();
  }

  const text = await response.text();
  return text as any;
}

/**
 * Check if request should include credentials
 */
export function shouldIncludeCredentials(url: string): boolean {
  return url.startsWith(API_CONFIG.BASE_URL);
}

/**
 * Build FormData from object
 */
export function buildFormData(data: Record<string, any>): FormData {
  const formData = new FormData();

  Object.entries(data).forEach(([key, value]) => {
    if (value !== undefined && value !== null) {
      if (value instanceof File) {
        formData.append(key, value);
      } else if (Array.isArray(value)) {
        value.forEach(item => formData.append(key, String(item)));
      } else {
        formData.append(key, String(value));
      }
    }
  });

  return formData;
}

/**
 * Download file from response
 */
export function downloadFile(blob: Blob, filename: string): void {
  const url = window.URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = filename;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  window.URL.revokeObjectURL(url);
}

/**
 * Abort controller with timeout
 */
export function createTimeoutController(timeout: number = API_CONFIG.TIMEOUT): AbortController {
  const controller = new AbortController();

  setTimeout(() => {
    controller.abort();
  }, timeout);

  return controller;
}

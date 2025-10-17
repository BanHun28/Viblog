/**
 * Error handling utilities for Viblog frontend
 * Standardized error responses and error code handling
 */

export interface ApiError {
  code: string;
  message: string;
  details?: any;
}

export interface ErrorResponse {
  error: ApiError;
}

/**
 * Error codes based on backend specification
 */
export enum ErrorCode {
  // Authentication errors
  UNAUTHORIZED = 'UNAUTHORIZED',
  FORBIDDEN = 'FORBIDDEN',
  INVALID_CREDENTIALS = 'INVALID_CREDENTIALS',
  TOKEN_EXPIRED = 'TOKEN_EXPIRED',

  // Validation errors
  VALIDATION_ERROR = 'VALIDATION_ERROR',
  INVALID_INPUT = 'INVALID_INPUT',

  // Resource errors
  NOT_FOUND = 'NOT_FOUND',
  ALREADY_EXISTS = 'ALREADY_EXISTS',

  // Rate limiting
  RATE_LIMIT_EXCEEDED = 'RATE_LIMIT_EXCEEDED',

  // Server errors
  INTERNAL_ERROR = 'INTERNAL_ERROR',
  SERVICE_UNAVAILABLE = 'SERVICE_UNAVAILABLE'
}

/**
 * Parse API error response
 */
export function parseApiError(error: any): ApiError {
  if (error.response?.data?.error) {
    return error.response.data.error;
  }

  if (error.response?.data) {
    return {
      code: ErrorCode.INTERNAL_ERROR,
      message: error.response.data.message || '알 수 없는 오류가 발생했습니다'
    };
  }

  if (error.message) {
    return {
      code: ErrorCode.INTERNAL_ERROR,
      message: error.message
    };
  }

  return {
    code: ErrorCode.INTERNAL_ERROR,
    message: '네트워크 오류가 발생했습니다'
  };
}

/**
 * Get user-friendly error message in Korean
 */
export function getErrorMessage(error: ApiError): string {
  const errorMessages: Record<string, string> = {
    [ErrorCode.UNAUTHORIZED]: '로그인이 필요합니다',
    [ErrorCode.FORBIDDEN]: '접근 권한이 없습니다',
    [ErrorCode.INVALID_CREDENTIALS]: '이메일 또는 비밀번호가 올바르지 않습니다',
    [ErrorCode.TOKEN_EXPIRED]: '로그인 세션이 만료되었습니다',
    [ErrorCode.VALIDATION_ERROR]: '입력 값을 확인해주세요',
    [ErrorCode.INVALID_INPUT]: '올바르지 않은 입력입니다',
    [ErrorCode.NOT_FOUND]: '요청한 리소스를 찾을 수 없습니다',
    [ErrorCode.ALREADY_EXISTS]: '이미 존재하는 항목입니다',
    [ErrorCode.RATE_LIMIT_EXCEEDED]: '요청 횟수가 초과되었습니다. 잠시 후 다시 시도해주세요',
    [ErrorCode.INTERNAL_ERROR]: '서버 오류가 발생했습니다',
    [ErrorCode.SERVICE_UNAVAILABLE]: '서비스를 일시적으로 사용할 수 없습니다'
  };

  return errorMessages[error.code] || error.message || '알 수 없는 오류가 발생했습니다';
}

/**
 * Check if error is authentication-related
 */
export function isAuthError(error: ApiError): boolean {
  return [
    ErrorCode.UNAUTHORIZED,
    ErrorCode.FORBIDDEN,
    ErrorCode.INVALID_CREDENTIALS,
    ErrorCode.TOKEN_EXPIRED
  ].includes(error.code as ErrorCode);
}

/**
 * Check if error is retryable
 */
export function isRetryableError(error: ApiError): boolean {
  return [
    ErrorCode.SERVICE_UNAVAILABLE,
    ErrorCode.RATE_LIMIT_EXCEEDED
  ].includes(error.code as ErrorCode);
}

/**
 * Handle API error with optional callback
 */
export function handleApiError(
  error: any,
  options?: {
    onAuthError?: () => void;
    onRetryableError?: () => void;
    showNotification?: (message: string) => void;
  }
): void {
  const apiError = parseApiError(error);
  const message = getErrorMessage(apiError);

  if (isAuthError(apiError) && options?.onAuthError) {
    options.onAuthError();
  } else if (isRetryableError(apiError) && options?.onRetryableError) {
    options.onRetryableError();
  }

  if (options?.showNotification) {
    options.showNotification(message);
  } else {
    console.error('API Error:', apiError);
  }
}

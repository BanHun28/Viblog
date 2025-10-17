/**
 * Local storage utilities for Viblog frontend
 * Safe localStorage wrapper with error handling
 */

/**
 * Storage keys used in the application
 */
export const StorageKeys = {
  AUTH_TOKEN: 'viblog_auth_token',
  REFRESH_TOKEN: 'viblog_refresh_token',
  USER_PREFERENCES: 'viblog_user_preferences',
  DRAFT_POST: 'viblog_draft_post',
  THEME: 'viblog_theme'
} as const;

/**
 * Check if localStorage is available
 */
export function isLocalStorageAvailable(): boolean {
  try {
    const test = '__localStorage_test__';
    localStorage.setItem(test, test);
    localStorage.removeItem(test);
    return true;
  } catch {
    return false;
  }
}

/**
 * Set item in localStorage with error handling
 */
export function setLocalStorage(key: string, value: any): boolean {
  if (!isLocalStorageAvailable()) {
    console.warn('localStorage is not available');
    return false;
  }

  try {
    const serialized = JSON.stringify(value);
    localStorage.setItem(key, serialized);
    return true;
  } catch (error) {
    console.error('Failed to set localStorage:', error);
    return false;
  }
}

/**
 * Get item from localStorage with error handling
 */
export function getLocalStorage<T>(key: string, defaultValue?: T): T | null {
  if (!isLocalStorageAvailable()) {
    return defaultValue || null;
  }

  try {
    const item = localStorage.getItem(key);
    if (item === null) {
      return defaultValue || null;
    }
    return JSON.parse(item) as T;
  } catch (error) {
    console.error('Failed to get localStorage:', error);
    return defaultValue || null;
  }
}

/**
 * Remove item from localStorage
 */
export function removeLocalStorage(key: string): boolean {
  if (!isLocalStorageAvailable()) {
    return false;
  }

  try {
    localStorage.removeItem(key);
    return true;
  } catch (error) {
    console.error('Failed to remove localStorage:', error);
    return false;
  }
}

/**
 * Clear all localStorage items
 */
export function clearLocalStorage(): boolean {
  if (!isLocalStorageAvailable()) {
    return false;
  }

  try {
    localStorage.clear();
    return true;
  } catch (error) {
    console.error('Failed to clear localStorage:', error);
    return false;
  }
}

/**
 * Get all keys from localStorage
 */
export function getLocalStorageKeys(): string[] {
  if (!isLocalStorageAvailable()) {
    return [];
  }

  try {
    return Object.keys(localStorage);
  } catch (error) {
    console.error('Failed to get localStorage keys:', error);
    return [];
  }
}

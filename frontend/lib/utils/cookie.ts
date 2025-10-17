/**
 * Cookie utilities for Viblog frontend
 * Safe cookie management with TypeScript types
 */

export interface CookieOptions {
  expires?: Date | number; // Date or days
  path?: string;
  domain?: string;
  secure?: boolean;
  sameSite?: 'strict' | 'lax' | 'none';
}

/**
 * Set cookie
 */
export function setCookie(
  name: string,
  value: string,
  options: CookieOptions = {}
): void {
  if (typeof document === 'undefined') {
    return;
  }

  let cookieString = `${encodeURIComponent(name)}=${encodeURIComponent(value)}`;

  // Expires
  if (options.expires) {
    let expires: Date;
    if (typeof options.expires === 'number') {
      expires = new Date();
      expires.setTime(expires.getTime() + options.expires * 24 * 60 * 60 * 1000);
    } else {
      expires = options.expires;
    }
    cookieString += `; expires=${expires.toUTCString()}`;
  }

  // Path
  if (options.path) {
    cookieString += `; path=${options.path}`;
  } else {
    cookieString += '; path=/';
  }

  // Domain
  if (options.domain) {
    cookieString += `; domain=${options.domain}`;
  }

  // Secure
  if (options.secure) {
    cookieString += '; secure';
  }

  // SameSite
  if (options.sameSite) {
    cookieString += `; samesite=${options.sameSite}`;
  }

  document.cookie = cookieString;
}

/**
 * Get cookie
 */
export function getCookie(name: string): string | null {
  if (typeof document === 'undefined') {
    return null;
  }

  const nameEQ = encodeURIComponent(name) + '=';
  const cookies = document.cookie.split(';');

  for (let cookie of cookies) {
    cookie = cookie.trim();
    if (cookie.indexOf(nameEQ) === 0) {
      return decodeURIComponent(cookie.substring(nameEQ.length));
    }
  }

  return null;
}

/**
 * Delete cookie
 */
export function deleteCookie(
  name: string,
  options: Pick<CookieOptions, 'path' | 'domain'> = {}
): void {
  setCookie(name, '', {
    ...options,
    expires: new Date(0)
  });
}

/**
 * Check if cookie exists
 */
export function hasCookie(name: string): boolean {
  return getCookie(name) !== null;
}

/**
 * Get all cookies as object
 */
export function getAllCookies(): Record<string, string> {
  if (typeof document === 'undefined') {
    return {};
  }

  const cookies: Record<string, string> = {};
  const cookieArray = document.cookie.split(';');

  for (let cookie of cookieArray) {
    cookie = cookie.trim();
    const [name, value] = cookie.split('=');
    if (name && value) {
      cookies[decodeURIComponent(name)] = decodeURIComponent(value);
    }
  }

  return cookies;
}

/**
 * Set JSON cookie
 */
export function setJsonCookie(
  name: string,
  value: any,
  options?: CookieOptions
): void {
  try {
    const jsonString = JSON.stringify(value);
    setCookie(name, jsonString, options);
  } catch (error) {
    console.error('Failed to set JSON cookie:', error);
  }
}

/**
 * Get JSON cookie
 */
export function getJsonCookie<T = any>(name: string): T | null {
  const value = getCookie(name);
  if (!value) {
    return null;
  }

  try {
    return JSON.parse(value) as T;
  } catch (error) {
    console.error('Failed to parse JSON cookie:', error);
    return null;
  }
}

/**
 * Clear all cookies
 */
export function clearAllCookies(): void {
  const cookies = getAllCookies();
  Object.keys(cookies).forEach(name => {
    deleteCookie(name);
  });
}

/**
 * Check if cookies are enabled
 */
export function areCookiesEnabled(): boolean {
  if (typeof document === 'undefined') {
    return false;
  }

  try {
    const testCookie = '__cookie_test__';
    setCookie(testCookie, 'test');
    const enabled = hasCookie(testCookie);
    deleteCookie(testCookie);
    return enabled;
  } catch {
    return false;
  }
}

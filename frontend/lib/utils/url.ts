/**
 * URL utilities for Viblog frontend
 * URL parsing, query string handling, path manipulation
 */

/**
 * Parse query string to object
 */
export function parseQueryString(queryString: string): Record<string, string> {
  const params = new URLSearchParams(queryString);
  const result: Record<string, string> = {};

  params.forEach((value, key) => {
    result[key] = value;
  });

  return result;
}

/**
 * Build query string from object
 */
export function buildQueryString(params: Record<string, any>): string {
  const searchParams = new URLSearchParams();

  Object.entries(params).forEach(([key, value]) => {
    if (value !== undefined && value !== null && value !== '') {
      if (Array.isArray(value)) {
        value.forEach(v => searchParams.append(key, String(v)));
      } else {
        searchParams.append(key, String(value));
      }
    }
  });

  const queryString = searchParams.toString();
  return queryString ? `?${queryString}` : '';
}

/**
 * Add or update query parameter
 */
export function addQueryParam(url: string, key: string, value: string): string {
  const urlObj = new URL(url, window.location.origin);
  urlObj.searchParams.set(key, value);
  return urlObj.toString();
}

/**
 * Remove query parameter
 */
export function removeQueryParam(url: string, key: string): string {
  const urlObj = new URL(url, window.location.origin);
  urlObj.searchParams.delete(key);
  return urlObj.toString();
}

/**
 * Get query parameter value
 */
export function getQueryParam(key: string): string | null {
  if (typeof window === 'undefined') {
    return null;
  }

  const params = new URLSearchParams(window.location.search);
  return params.get(key);
}

/**
 * Get all query parameters
 */
export function getAllQueryParams(): Record<string, string> {
  if (typeof window === 'undefined') {
    return {};
  }

  return parseQueryString(window.location.search);
}

/**
 * Check if URL is absolute
 */
export function isAbsoluteUrl(url: string): boolean {
  return /^https?:\/\//i.test(url);
}

/**
 * Check if URL is external
 */
export function isExternalUrl(url: string): boolean {
  if (!isAbsoluteUrl(url)) {
    return false;
  }

  try {
    const urlObj = new URL(url);
    return urlObj.origin !== window.location.origin;
  } catch {
    return false;
  }
}

/**
 * Join URL paths
 */
export function joinPaths(...paths: string[]): string {
  return paths
    .map((path, index) => {
      if (index === 0) {
        return path.replace(/\/+$/, '');
      }
      return path.replace(/^\/+/, '').replace(/\/+$/, '');
    })
    .filter(Boolean)
    .join('/');
}

/**
 * Get base URL (origin)
 */
export function getBaseUrl(): string {
  if (typeof window === 'undefined') {
    return process.env.NEXT_PUBLIC_SITE_URL || 'http://localhost:30001';
  }
  return window.location.origin;
}

/**
 * Build full URL from path
 */
export function buildFullUrl(path: string): string {
  const baseUrl = getBaseUrl();
  const cleanPath = path.startsWith('/') ? path : `/${path}`;
  return `${baseUrl}${cleanPath}`;
}

/**
 * Extract path from URL
 */
export function extractPath(url: string): string {
  try {
    const urlObj = new URL(url);
    return urlObj.pathname + urlObj.search + urlObj.hash;
  } catch {
    return url;
  }
}

/**
 * Normalize URL (remove trailing slash, lowercase)
 */
export function normalizeUrl(url: string): string {
  return url.replace(/\/+$/, '').toLowerCase();
}

/**
 * Get URL without query string and hash
 */
export function getCleanUrl(url: string): string {
  try {
    const urlObj = new URL(url, window.location.origin);
    return `${urlObj.origin}${urlObj.pathname}`;
  } catch {
    return url.split('?')[0].split('#')[0];
  }
}

/**
 * Update current URL without page reload
 */
export function updateUrlWithoutReload(url: string): void {
  if (typeof window === 'undefined') {
    return;
  }

  window.history.pushState({}, '', url);
}

/**
 * Replace current URL without page reload
 */
export function replaceUrlWithoutReload(url: string): void {
  if (typeof window === 'undefined') {
    return;
  }

  window.history.replaceState({}, '', url);
}

/**
 * Build post URL
 */
export function buildPostUrl(postId: string): string {
  return `/posts/${postId}`;
}

/**
 * Build category URL
 */
export function buildCategoryUrl(categorySlug: string): string {
  return `/categories/${categorySlug}`;
}

/**
 * Build tag URL
 */
export function buildTagUrl(tagSlug: string): string {
  return `/tags/${tagSlug}`;
}

/**
 * Build search URL
 */
export function buildSearchUrl(query: string): string {
  return `/search?q=${encodeURIComponent(query)}`;
}

/**
 * Build author URL
 */
export function buildAuthorUrl(authorId: string): string {
  return `/authors/${authorId}`;
}

/**
 * Extract post ID from URL
 */
export function extractPostIdFromUrl(url: string): string | null {
  const match = url.match(/\/posts\/([^\/\?#]+)/);
  return match ? match[1] : null;
}

/**
 * Check if current page is active
 */
export function isActivePage(path: string): boolean {
  if (typeof window === 'undefined') {
    return false;
  }

  return window.location.pathname === path;
}

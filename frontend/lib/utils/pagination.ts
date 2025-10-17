/**
 * Pagination utilities for Viblog frontend
 * Supports cursor-based pagination (20 items per page)
 */

export interface PaginationParams {
  cursor?: string;
  limit?: number;
}

export interface PaginationInfo {
  hasMore: boolean;
  nextCursor?: string;
  totalCount?: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  pagination: PaginationInfo;
}

/**
 * Default pagination limit (페이지당 20개)
 */
export const DEFAULT_PAGE_LIMIT = 20;

/**
 * Build pagination query parameters
 */
export function buildPaginationParams(cursor?: string, limit: number = DEFAULT_PAGE_LIMIT): string {
  const params = new URLSearchParams();

  if (cursor) {
    params.append('cursor', cursor);
  }

  params.append('limit', limit.toString());

  return params.toString();
}

/**
 * Parse pagination response
 */
export function parsePaginationResponse<T>(response: any): PaginatedResponse<T> {
  return {
    data: response.data || [],
    pagination: {
      hasMore: response.has_more || false,
      nextCursor: response.next_cursor,
      totalCount: response.total_count
    }
  };
}

/**
 * Check if should load more items
 */
export function shouldLoadMore(
  isLoading: boolean,
  hasMore: boolean,
  element: HTMLElement | null
): boolean {
  if (!element || isLoading || !hasMore) {
    return false;
  }

  const rect = element.getBoundingClientRect();
  const viewHeight = window.innerHeight || document.documentElement.clientHeight;

  return rect.bottom <= viewHeight + 200; // 200px threshold
}

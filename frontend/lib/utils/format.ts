/**
 * Formatting utilities for Viblog frontend
 * Handles date, number, text formatting
 */

/**
 * Format date for display
 * Returns localized Korean date format
 */
export function formatDate(date: Date | string): string {
  const d = typeof date === 'string' ? new Date(date) : date;

  return new Intl.DateTimeFormat('ko-KR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  }).format(d);
}

/**
 * Format datetime for display
 * Returns localized Korean datetime format
 */
export function formatDateTime(date: Date | string): string {
  const d = typeof date === 'string' ? new Date(date) : date;

  return new Intl.DateTimeFormat('ko-KR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(d);
}

/**
 * Format relative time (e.g., "3시간 전", "2일 전")
 * For comment timestamps and post dates
 */
export function formatRelativeTime(date: Date | string): string {
  const d = typeof date === 'string' ? new Date(date) : date;
  const now = new Date();
  const diffInSeconds = Math.floor((now.getTime() - d.getTime()) / 1000);

  if (diffInSeconds < 60) {
    return '방금 전';
  }

  const diffInMinutes = Math.floor(diffInSeconds / 60);
  if (diffInMinutes < 60) {
    return `${diffInMinutes}분 전`;
  }

  const diffInHours = Math.floor(diffInMinutes / 60);
  if (diffInHours < 24) {
    return `${diffInHours}시간 전`;
  }

  const diffInDays = Math.floor(diffInHours / 24);
  if (diffInDays < 7) {
    return `${diffInDays}일 전`;
  }

  const diffInWeeks = Math.floor(diffInDays / 7);
  if (diffInWeeks < 4) {
    return `${diffInWeeks}주 전`;
  }

  const diffInMonths = Math.floor(diffInDays / 30);
  if (diffInMonths < 12) {
    return `${diffInMonths}개월 전`;
  }

  const diffInYears = Math.floor(diffInDays / 365);
  return `${diffInYears}년 전`;
}

/**
 * Format number with thousands separator
 * For view counts, like counts, etc.
 */
export function formatNumber(num: number): string {
  return new Intl.NumberFormat('ko-KR').format(num);
}

/**
 * Format large numbers with K, M suffix
 * For compact display of large counts
 */
export function formatCompactNumber(num: number): string {
  if (num < 1000) {
    return num.toString();
  }

  if (num < 1000000) {
    return `${(num / 1000).toFixed(1)}K`;
  }

  return `${(num / 1000000).toFixed(1)}M`;
}

/**
 * Truncate text with ellipsis
 * For previews and excerpts
 */
export function truncateText(text: string, maxLength: number): string {
  if (text.length <= maxLength) {
    return text;
  }

  return text.slice(0, maxLength).trim() + '...';
}

/**
 * Extract excerpt from markdown content
 * Removes markdown syntax and truncates
 */
export function extractExcerpt(content: string, maxLength: number = 200): string {
  // Remove markdown syntax
  let plainText = content
    .replace(/#{1,6}\s/g, '') // headers
    .replace(/\*\*(.+?)\*\*/g, '$1') // bold
    .replace(/\*(.+?)\*/g, '$1') // italic
    .replace(/\[(.+?)\]\(.+?\)/g, '$1') // links
    .replace(/`(.+?)`/g, '$1') // inline code
    .replace(/```[\s\S]*?```/g, '') // code blocks
    .replace(/!\[.*?\]\(.*?\)/g, '') // images
    .replace(/\n+/g, ' ') // newlines
    .trim();

  return truncateText(plainText, maxLength);
}

/**
 * Sanitize filename for safe use
 * Removes special characters and spaces
 */
export function sanitizeFilename(filename: string): string {
  return filename
    .replace(/[^a-zA-Z0-9가-힣.\-_]/g, '_')
    .replace(/_{2,}/g, '_')
    .trim();
}

/**
 * Format file size in human-readable format
 * For file upload displays
 */
export function formatFileSize(bytes: number): string {
  if (bytes === 0) return '0 Bytes';

  const k = 1024;
  const sizes = ['Bytes', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`;
}

/**
 * Highlight search term in text
 * Returns HTML string with highlighted terms
 */
export function highlightSearchTerm(text: string, searchTerm: string): string {
  if (!searchTerm) return text;

  const regex = new RegExp(`(${searchTerm})`, 'gi');
  return text.replace(regex, '<mark>$1</mark>');
}

/**
 * Generate slug from title
 * For URL-friendly post slugs
 */
export function generateSlug(title: string): string {
  return title
    .toLowerCase()
    .trim()
    .replace(/[^\w\s가-힣-]/g, '')
    .replace(/[\s_-]+/g, '-')
    .replace(/^-+|-+$/g, '');
}

/**
 * Parse tags from comma-separated string
 * Returns array of trimmed, unique tags
 */
export function parseTags(tagsString: string): string[] {
  return tagsString
    .split(',')
    .map(tag => tag.trim())
    .filter(tag => tag.length > 0)
    .filter((tag, index, self) => self.indexOf(tag) === index);
}

/**
 * Format tags to comma-separated string
 */
export function formatTags(tags: string[]): string {
  return tags.join(', ');
}

/**
 * Calculate reading time in minutes
 * Based on average reading speed of 200 words per minute
 */
export function calculateReadingTime(content: string): number {
  const wordsPerMinute = 200;
  const words = content.trim().split(/\s+/).length;
  const minutes = Math.ceil(words / wordsPerMinute);
  return Math.max(1, minutes);
}

/**
 * Format reading time for display
 */
export function formatReadingTime(minutes: number): string {
  if (minutes < 1) return '1분';
  return `${minutes}분`;
}

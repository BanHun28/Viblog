/**
 * Constants for Viblog frontend
 * Based on product specifications
 */

/**
 * API Configuration
 */
export const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:30000';
export const API_VERSION = 'v1';
export const API_TIMEOUT = 30000; // 30 seconds

/**
 * Pagination Configuration (페이지당 20개)
 */
export const POSTS_PER_PAGE = 20;
export const COMMENTS_PER_PAGE = 20;

/**
 * Validation Limits
 */
export const VALIDATION_LIMITS = {
  // User
  EMAIL_MAX_LENGTH: 255,
  PASSWORD_MIN_LENGTH: 8,
  PASSWORD_MAX_LENGTH: 128,
  NICKNAME_MIN_LENGTH: 2,
  NICKNAME_MAX_LENGTH: 20,

  // Post
  POST_TITLE_MAX_LENGTH: 200,
  POST_CONTENT_MAX_LENGTH: 100000,
  POST_EXCERPT_MAX_LENGTH: 500,
  TAGS_MAX_COUNT: 10, // 글당 최대 10개
  TAG_MAX_LENGTH: 30,

  // Comment
  COMMENT_MIN_LENGTH: 2,
  COMMENT_MAX_LENGTH: 1000,
  COMMENT_NICKNAME_MIN_LENGTH: 2,
  COMMENT_NICKNAME_MAX_LENGTH: 20,
  COMMENT_PASSWORD_MIN_LENGTH: 4,

  // Category
  CATEGORY_NAME_MAX_LENGTH: 50,
  CATEGORY_SLUG_MAX_LENGTH: 100
} as const;

/**
 * JWT Token Configuration
 */
export const TOKEN_CONFIG = {
  ACCESS_TOKEN_DURATION: 15 * 60, // 15분 (초 단위)
  REFRESH_TOKEN_DURATION: 7 * 24 * 60 * 60, // 7일 (초 단위)
  TOKEN_REFRESH_THRESHOLD: 5 * 60 // 5분 전 갱신
} as const;

/**
 * Rate Limiting Configuration
 */
export const RATE_LIMITS = {
  API_REQUESTS_PER_MINUTE: 100,
  COMMENTS_PER_MINUTE: 5,
  LOGIN_ATTEMPTS_PER_HOUR: 5
} as const;

/**
 * Post Status (글 상태)
 */
export enum PostStatus {
  DRAFT = 'draft', // 임시저장
  PUBLISHED = 'published', // 발행
  SCHEDULED = 'scheduled' // 예약발행
}

/**
 * User Roles (권한 체계)
 */
export enum UserRole {
  ADMIN = 'admin', // 관리자
  USER = 'user' // 일반 사용자
}

/**
 * Comment Type
 */
export enum CommentType {
  MEMBER = 'member', // 회원 댓글
  ANONYMOUS = 'anonymous' // 익명 댓글
}

/**
 * Notification Types
 */
export enum NotificationType {
  REPLY = 'reply', // 답글 알림
  LIKE = 'like', // 좋아요 알림
  MENTION = 'mention' // 멘션 알림
}

/**
 * Sort Options
 */
export const SORT_OPTIONS = {
  LATEST: 'latest', // 최신순
  OLDEST: 'oldest', // 오래된순
  POPULAR: 'popular', // 인기순 (조회수 + 좋아요)
  MOST_VIEWED: 'most_viewed', // 조회수순
  MOST_LIKED: 'most_liked' // 좋아요순
} as const;

/**
 * Search Configuration
 */
export const SEARCH_CONFIG = {
  MIN_QUERY_LENGTH: 2,
  MAX_QUERY_LENGTH: 100,
  DEBOUNCE_DELAY: 300, // ms
  HIGHLIGHT_CLASS: 'search-highlight'
} as const;

/**
 * Image Configuration
 */
export const IMAGE_CONFIG = {
  MAX_SIZE: 5 * 1024 * 1024, // 5MB
  ALLOWED_TYPES: ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp'],
  DEFAULT_AVATAR: '/images/default-avatar.png'
} as const;

/**
 * Cache Configuration
 */
export const CACHE_CONFIG = {
  VIEW_COUNT_DURATION: 24 * 60 * 60, // 24시간 (초 단위)
  POST_CACHE_DURATION: 5 * 60, // 5분 (SSG/ISR)
  API_CACHE_DURATION: 60 // 1분
} as const;

/**
 * UI Configuration
 */
export const UI_CONFIG = {
  TOAST_DURATION: 3000, // 3 seconds
  MODAL_ANIMATION_DURATION: 200, // ms
  DEBOUNCE_DELAY: 300, // ms
  THROTTLE_DELAY: 1000, // ms
  INFINITE_SCROLL_THRESHOLD: 200 // px from bottom
} as const;

/**
 * Localization
 */
export const SUPPORTED_LANGUAGES = ['ko', 'en'] as const;
export const DEFAULT_LANGUAGE = 'ko';

/**
 * Date Format
 */
export const DATE_FORMAT = {
  SHORT: 'YYYY-MM-DD',
  LONG: 'YYYY년 MM월 DD일',
  DATETIME: 'YYYY-MM-DD HH:mm',
  TIME: 'HH:mm'
} as const;

/**
 * Regular Expressions
 */
export const REGEX = {
  EMAIL: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
  URL: /^https?:\/\/.+/,
  SLUG: /^[a-z0-9-]+$/,
  PASSWORD: /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/
} as const;

/**
 * Breakpoints for responsive design
 */
export const BREAKPOINTS = {
  MOBILE: 640, // px
  TABLET: 768, // px
  DESKTOP: 1024, // px
  WIDE: 1280 // px
} as const;

/**
 * Z-Index layers
 */
export const Z_INDEX = {
  DROPDOWN: 1000,
  MODAL: 1050,
  TOAST: 1100,
  TOOLTIP: 1200
} as const;

/**
 * External links
 */
export const EXTERNAL_LINKS = {
  GITHUB: 'https://github.com/yourusername/viblog',
  EMAIL: 'contact@viblog.com'
} as const;

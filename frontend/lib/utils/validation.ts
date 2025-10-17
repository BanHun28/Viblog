/**
 * Validation utilities for Viblog frontend
 * Based on product specifications requirements
 */

/**
 * Email validation
 * Required for user registration and login
 */
export function validateEmail(email: string): boolean {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
}

/**
 * Password validation
 * Requirements: 최소 8자, 영문+숫자+특수문자
 */
export function validatePassword(password: string): {
  isValid: boolean;
  errors: string[];
} {
  const errors: string[] = [];

  if (password.length < 8) {
    errors.push('비밀번호는 최소 8자 이상이어야 합니다');
  }

  if (!/[a-zA-Z]/.test(password)) {
    errors.push('영문자를 포함해야 합니다');
  }

  if (!/[0-9]/.test(password)) {
    errors.push('숫자를 포함해야 합니다');
  }

  if (!/[!@#$%^&*(),.?":{}|<>]/.test(password)) {
    errors.push('특수문자를 포함해야 합니다');
  }

  return {
    isValid: errors.length === 0,
    errors
  };
}

/**
 * Nickname validation
 * Required for user registration
 */
export function validateNickname(nickname: string): {
  isValid: boolean;
  error?: string;
} {
  if (!nickname || nickname.trim().length === 0) {
    return { isValid: false, error: '닉네임을 입력해주세요' };
  }

  if (nickname.length < 2) {
    return { isValid: false, error: '닉네임은 최소 2자 이상이어야 합니다' };
  }

  if (nickname.length > 20) {
    return { isValid: false, error: '닉네임은 최대 20자까지 가능합니다' };
  }

  return { isValid: true };
}

/**
 * URL validation
 * For profile images and external links
 */
export function validateUrl(url: string): boolean {
  try {
    new URL(url);
    return true;
  } catch {
    return false;
  }
}

/**
 * Tag validation
 * Maximum 10 tags per post
 */
export function validateTags(tags: string[]): {
  isValid: boolean;
  error?: string;
} {
  if (tags.length > 10) {
    return { isValid: false, error: '태그는 최대 10개까지 가능합니다' };
  }

  for (const tag of tags) {
    if (tag.trim().length === 0) {
      return { isValid: false, error: '빈 태그는 허용되지 않습니다' };
    }

    if (tag.length > 30) {
      return { isValid: false, error: '태그는 최대 30자까지 가능합니다' };
    }
  }

  return { isValid: true };
}

/**
 * Post title validation
 */
export function validatePostTitle(title: string): {
  isValid: boolean;
  error?: string;
} {
  if (!title || title.trim().length === 0) {
    return { isValid: false, error: '제목을 입력해주세요' };
  }

  if (title.length > 200) {
    return { isValid: false, error: '제목은 최대 200자까지 가능합니다' };
  }

  return { isValid: true };
}

/**
 * Comment content validation
 */
export function validateComment(content: string): {
  isValid: boolean;
  error?: string;
} {
  if (!content || content.trim().length === 0) {
    return { isValid: false, error: '댓글 내용을 입력해주세요' };
  }

  if (content.length < 2) {
    return { isValid: false, error: '댓글은 최소 2자 이상이어야 합니다' };
  }

  if (content.length > 1000) {
    return { isValid: false, error: '댓글은 최대 1000자까지 가능합니다' };
  }

  return { isValid: true };
}

/**
 * Anonymous comment validation
 * For nickname and password
 */
export function validateAnonymousComment(nickname: string, password: string): {
  isValid: boolean;
  errors: { nickname?: string; password?: string };
} {
  const errors: { nickname?: string; password?: string } = {};

  if (!nickname || nickname.trim().length === 0) {
    errors.nickname = '닉네임을 입력해주세요';
  } else if (nickname.length < 2 || nickname.length > 20) {
    errors.nickname = '닉네임은 2-20자 사이여야 합니다';
  }

  if (!password || password.length < 4) {
    errors.password = '비밀번호는 최소 4자 이상이어야 합니다';
  }

  return {
    isValid: Object.keys(errors).length === 0,
    errors
  };
}

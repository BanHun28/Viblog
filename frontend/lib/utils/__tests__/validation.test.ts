/**
 * Validation utilities tests
 */

import { describe, it, expect } from 'vitest';
import {
  validateEmail,
  validatePassword,
  validateNickname,
  validateUrl,
  validateTags,
  validatePostTitle,
  validateComment,
  validateAnonymousComment
} from '../validation';

describe('validateEmail', () => {
  it('유효한 이메일 주소를 검증해야 함', () => {
    expect(validateEmail('test@example.com')).toBe(true);
    expect(validateEmail('user.name@domain.co.kr')).toBe(true);
    expect(validateEmail('test+tag@gmail.com')).toBe(true);
  });

  it('잘못된 이메일 주소를 거부해야 함', () => {
    expect(validateEmail('invalid')).toBe(false);
    expect(validateEmail('test@')).toBe(false);
    expect(validateEmail('@domain.com')).toBe(false);
    expect(validateEmail('test @domain.com')).toBe(false);
  });
});

describe('validatePassword', () => {
  it('유효한 비밀번호를 검증해야 함', () => {
    const result = validatePassword('Test1234!');
    expect(result.isValid).toBe(true);
    expect(result.errors).toHaveLength(0);
  });

  it('8자 미만 비밀번호를 거부해야 함', () => {
    const result = validatePassword('Test1!');
    expect(result.isValid).toBe(false);
    expect(result.errors).toContain('비밀번호는 최소 8자 이상이어야 합니다');
  });

  it('영문자가 없는 비밀번호를 거부해야 함', () => {
    const result = validatePassword('12345678!');
    expect(result.isValid).toBe(false);
    expect(result.errors).toContain('영문자를 포함해야 합니다');
  });

  it('숫자가 없는 비밀번호를 거부해야 함', () => {
    const result = validatePassword('TestTest!');
    expect(result.isValid).toBe(false);
    expect(result.errors).toContain('숫자를 포함해야 합니다');
  });

  it('특수문자가 없는 비밀번호를 거부해야 함', () => {
    const result = validatePassword('Test1234');
    expect(result.isValid).toBe(false);
    expect(result.errors).toContain('특수문자를 포함해야 합니다');
  });
});

describe('validateNickname', () => {
  it('유효한 닉네임을 검증해야 함', () => {
    expect(validateNickname('테스터').isValid).toBe(true);
    expect(validateNickname('User123').isValid).toBe(true);
    expect(validateNickname('닉네임테스트').isValid).toBe(true);
  });

  it('빈 닉네임을 거부해야 함', () => {
    const result = validateNickname('');
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('닉네임을 입력해주세요');
  });

  it('2자 미만 닉네임을 거부해야 함', () => {
    const result = validateNickname('A');
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('닉네임은 최소 2자 이상이어야 합니다');
  });

  it('20자 초과 닉네임을 거부해야 함', () => {
    const result = validateNickname('A'.repeat(21));
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('닉네임은 최대 20자까지 가능합니다');
  });
});

describe('validateUrl', () => {
  it('유효한 URL을 검증해야 함', () => {
    expect(validateUrl('https://example.com')).toBe(true);
    expect(validateUrl('http://localhost:3000')).toBe(true);
    expect(validateUrl('https://example.com/path/to/page')).toBe(true);
  });

  it('잘못된 URL을 거부해야 함', () => {
    expect(validateUrl('not a url')).toBe(false);
    expect(validateUrl('example.com')).toBe(false);
  });
});

describe('validateTags', () => {
  it('유효한 태그 배열을 검증해야 함', () => {
    expect(validateTags(['javascript', 'typescript']).isValid).toBe(true);
    expect(validateTags(['tag1', 'tag2', 'tag3']).isValid).toBe(true);
  });

  it('10개 초과 태그를 거부해야 함', () => {
    const tags = Array(11).fill('tag');
    const result = validateTags(tags);
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('태그는 최대 10개까지 가능합니다');
  });

  it('빈 태그를 거부해야 함', () => {
    const result = validateTags(['tag1', '', 'tag3']);
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('빈 태그는 허용되지 않습니다');
  });

  it('30자 초과 태그를 거부해야 함', () => {
    const result = validateTags(['a'.repeat(31)]);
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('태그는 최대 30자까지 가능합니다');
  });
});

describe('validatePostTitle', () => {
  it('유효한 제목을 검증해야 함', () => {
    expect(validatePostTitle('테스트 포스트').isValid).toBe(true);
    expect(validatePostTitle('A'.repeat(200)).isValid).toBe(true);
  });

  it('빈 제목을 거부해야 함', () => {
    const result = validatePostTitle('');
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('제목을 입력해주세요');
  });

  it('200자 초과 제목을 거부해야 함', () => {
    const result = validatePostTitle('A'.repeat(201));
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('제목은 최대 200자까지 가능합니다');
  });
});

describe('validateComment', () => {
  it('유효한 댓글을 검증해야 함', () => {
    expect(validateComment('좋은 글이네요!').isValid).toBe(true);
    expect(validateComment('테스트 댓글입니다.').isValid).toBe(true);
  });

  it('빈 댓글을 거부해야 함', () => {
    const result = validateComment('');
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('댓글 내용을 입력해주세요');
  });

  it('2자 미만 댓글을 거부해야 함', () => {
    const result = validateComment('A');
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('댓글은 최소 2자 이상이어야 합니다');
  });

  it('1000자 초과 댓글을 거부해야 함', () => {
    const result = validateComment('A'.repeat(1001));
    expect(result.isValid).toBe(false);
    expect(result.error).toBe('댓글은 최대 1000자까지 가능합니다');
  });
});

describe('validateAnonymousComment', () => {
  it('유효한 익명 댓글 정보를 검증해야 함', () => {
    const result = validateAnonymousComment('익명', '1234');
    expect(result.isValid).toBe(true);
    expect(result.errors).toEqual({});
  });

  it('빈 닉네임을 거부해야 함', () => {
    const result = validateAnonymousComment('', '1234');
    expect(result.isValid).toBe(false);
    expect(result.errors.nickname).toBe('닉네임을 입력해주세요');
  });

  it('짧은 비밀번호를 거부해야 함', () => {
    const result = validateAnonymousComment('익명', '123');
    expect(result.isValid).toBe(false);
    expect(result.errors.password).toBe('비밀번호는 최소 4자 이상이어야 합니다');
  });

  it('여러 검증 오류를 동시에 반환해야 함', () => {
    const result = validateAnonymousComment('', '12');
    expect(result.isValid).toBe(false);
    expect(result.errors.nickname).toBeDefined();
    expect(result.errors.password).toBeDefined();
  });
});

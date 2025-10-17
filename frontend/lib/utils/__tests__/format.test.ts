/**
 * Formatting utilities tests
 */

import { describe, it, expect } from 'vitest';
import {
  formatDate,
  formatDateTime,
  formatRelativeTime,
  formatNumber,
  formatCompactNumber,
  truncateText,
  extractExcerpt,
  sanitizeFilename,
  formatFileSize,
  generateSlug,
  parseTags,
  formatTags,
  calculateReadingTime,
  formatReadingTime
} from '../format';

describe('formatDate', () => {
  it('날짜를 한국어 형식으로 포맷팅해야 함', () => {
    const date = new Date('2024-01-15');
    const formatted = formatDate(date);
    expect(formatted).toContain('2024');
    expect(formatted).toContain('1');
    expect(formatted).toContain('15');
  });

  it('문자열 날짜도 포맷팅해야 함', () => {
    const formatted = formatDate('2024-01-15');
    expect(formatted).toContain('2024');
  });
});

describe('formatRelativeTime', () => {
  it('방금 전을 반환해야 함', () => {
    const now = new Date();
    expect(formatRelativeTime(now)).toBe('방금 전');
  });

  it('분 단위를 반환해야 함', () => {
    const date = new Date(Date.now() - 5 * 60 * 1000); // 5분 전
    expect(formatRelativeTime(date)).toBe('5분 전');
  });

  it('시간 단위를 반환해야 함', () => {
    const date = new Date(Date.now() - 3 * 60 * 60 * 1000); // 3시간 전
    expect(formatRelativeTime(date)).toBe('3시간 전');
  });

  it('일 단위를 반환해야 함', () => {
    const date = new Date(Date.now() - 2 * 24 * 60 * 60 * 1000); // 2일 전
    expect(formatRelativeTime(date)).toBe('2일 전');
  });
});

describe('formatNumber', () => {
  it('천 단위 구분자를 추가해야 함', () => {
    expect(formatNumber(1000)).toContain('1,000');
    expect(formatNumber(1000000)).toContain('1,000,000');
  });

  it('작은 숫자는 그대로 반환해야 함', () => {
    expect(formatNumber(100)).toBe('100');
  });
});

describe('formatCompactNumber', () => {
  it('1000 미만은 그대로 반환해야 함', () => {
    expect(formatCompactNumber(999)).toBe('999');
  });

  it('천 단위는 K로 표시해야 함', () => {
    expect(formatCompactNumber(1500)).toBe('1.5K');
    expect(formatCompactNumber(10000)).toBe('10.0K');
  });

  it('백만 단위는 M으로 표시해야 함', () => {
    expect(formatCompactNumber(1500000)).toBe('1.5M');
  });
});

describe('truncateText', () => {
  it('최대 길이보다 짧은 텍스트는 그대로 반환해야 함', () => {
    expect(truncateText('짧은 텍스트', 20)).toBe('짧은 텍스트');
  });

  it('긴 텍스트는 잘라야 함', () => {
    const text = '긴 텍스트입니다. 이 텍스트는 잘려야 합니다.';
    const truncated = truncateText(text, 10);
    expect(truncated).toHaveLength(13); // 10 + '...'
    expect(truncated).toContain('...');
  });
});

describe('extractExcerpt', () => {
  it('마크다운 문법을 제거해야 함', () => {
    const markdown = '# 제목\n\n**굵은 글씨**와 *기울임꼴*입니다.';
    const excerpt = extractExcerpt(markdown, 100);
    expect(excerpt).not.toContain('#');
    expect(excerpt).not.toContain('**');
    expect(excerpt).not.toContain('*');
  });

  it('코드 블록을 제거해야 함', () => {
    const markdown = '텍스트\n```js\nconst x = 1;\n```\n더 많은 텍스트';
    const excerpt = extractExcerpt(markdown, 100);
    expect(excerpt).not.toContain('```');
    expect(excerpt).not.toContain('const');
  });

  it('최대 길이로 잘라야 함', () => {
    const markdown = 'A'.repeat(300);
    const excerpt = extractExcerpt(markdown, 50);
    expect(excerpt.length).toBeLessThanOrEqual(53); // 50 + '...'
  });
});

describe('sanitizeFilename', () => {
  it('특수 문자를 언더스코어로 바꿔야 함', () => {
    expect(sanitizeFilename('file name!@#.txt')).toBe('file_name___.txt');
  });

  it('한글 파일명을 유지해야 함', () => {
    expect(sanitizeFilename('파일이름.txt')).toBe('파일이름.txt');
  });

  it('연속된 언더스코어를 하나로 합쳐야 함', () => {
    expect(sanitizeFilename('file___name.txt')).toBe('file_name.txt');
  });
});

describe('formatFileSize', () => {
  it('0 바이트를 처리해야 함', () => {
    expect(formatFileSize(0)).toBe('0 Bytes');
  });

  it('바이트 단위를 표시해야 함', () => {
    expect(formatFileSize(500)).toBe('500 Bytes');
  });

  it('킬로바이트 단위를 표시해야 함', () => {
    expect(formatFileSize(1024)).toBe('1 KB');
    expect(formatFileSize(1536)).toBe('1.5 KB');
  });

  it('메가바이트 단위를 표시해야 함', () => {
    expect(formatFileSize(1024 * 1024)).toBe('1 MB');
  });
});

describe('generateSlug', () => {
  it('공백을 하이픈으로 바꿔야 함', () => {
    expect(generateSlug('Hello World')).toBe('hello-world');
  });

  it('특수 문자를 제거해야 함', () => {
    expect(generateSlug('Hello! World?')).toBe('hello-world');
  });

  it('한글도 처리해야 함', () => {
    expect(generateSlug('안녕하세요 세계')).toBe('안녕하세요-세계');
  });

  it('연속된 하이픈을 하나로 합쳐야 함', () => {
    expect(generateSlug('Hello   World')).toBe('hello-world');
  });

  it('앞뒤 하이픈을 제거해야 함', () => {
    expect(generateSlug('-Hello World-')).toBe('hello-world');
  });
});

describe('parseTags', () => {
  it('쉼표로 구분된 태그를 배열로 변환해야 함', () => {
    const tags = parseTags('javascript, typescript, react');
    expect(tags).toEqual(['javascript', 'typescript', 'react']);
  });

  it('공백을 제거해야 함', () => {
    const tags = parseTags('  tag1  ,  tag2  ');
    expect(tags).toEqual(['tag1', 'tag2']);
  });

  it('빈 태그를 제거해야 함', () => {
    const tags = parseTags('tag1, , tag2');
    expect(tags).toEqual(['tag1', 'tag2']);
  });

  it('중복 태그를 제거해야 함', () => {
    const tags = parseTags('tag1, tag2, tag1');
    expect(tags).toEqual(['tag1', 'tag2']);
  });
});

describe('formatTags', () => {
  it('배열을 쉼표로 구분된 문자열로 변환해야 함', () => {
    expect(formatTags(['tag1', 'tag2', 'tag3'])).toBe('tag1, tag2, tag3');
  });

  it('빈 배열은 빈 문자열을 반환해야 함', () => {
    expect(formatTags([])).toBe('');
  });
});

describe('calculateReadingTime', () => {
  it('짧은 텍스트는 최소 1분을 반환해야 함', () => {
    expect(calculateReadingTime('짧은 텍스트')).toBe(1);
  });

  it('긴 텍스트는 단어 수에 비례해야 함', () => {
    const words = Array(400).fill('word').join(' ');
    const minutes = calculateReadingTime(words);
    expect(minutes).toBe(2); // 400 words / 200 wpm = 2 minutes
  });
});

describe('formatReadingTime', () => {
  it('분 단위를 한국어로 포맷팅해야 함', () => {
    expect(formatReadingTime(1)).toBe('1분');
    expect(formatReadingTime(5)).toBe('5분');
  });

  it('1분 미만은 1분으로 표시해야 함', () => {
    expect(formatReadingTime(0)).toBe('1분');
  });
});

# Viblog Frontend Utilities Test Suite

프론트엔드 유틸리티 함수에 대한 포괄적인 테스트 스위트입니다.

## 🧪 테스트 프레임워크

- **Vitest** - 빠르고 현대적인 테스트 프레임워크
- **@testing-library/react** - React 컴포넌트 테스팅
- **@testing-library/jest-dom** - DOM 매처 확장
- **happy-dom** - 빠른 DOM 구현체

## 📦 설치

```bash
npm install
```

## 🚀 테스트 실행

### 모든 테스트 실행
```bash
npm test
```

### Watch 모드로 실행
```bash
npm test -- --watch
```

### UI 모드로 실행
```bash
npm run test:ui
```

### 커버리지 리포트 생성
```bash
npm run test:coverage
```

## 📁 테스트 구조

```
lib/utils/__tests__/
├── validation.test.ts    # 검증 유틸리티 테스트
├── format.test.ts        # 포맷팅 유틸리티 테스트
├── collection.test.ts    # 컬렉션 유틸리티 테스트
├── url.test.ts          # URL 유틸리티 테스트 (추가 예정)
├── storage.test.ts      # 저장소 유틸리티 테스트 (추가 예정)
└── ...
```

## ✅ 테스트 범위

### validation.test.ts
- ✅ 이메일 검증
- ✅ 비밀번호 검증 (8자 이상, 영문+숫자+특수문자)
- ✅ 닉네임 검증 (2-20자)
- ✅ URL 검증
- ✅ 태그 검증 (최대 10개)
- ✅ 포스트 제목 검증
- ✅ 댓글 검증 (2-1000자)
- ✅ 익명 댓글 검증

### format.test.ts
- ✅ 날짜/시간 포맷팅 (한국어)
- ✅ 상대 시간 ("방금 전", "5분 전")
- ✅ 숫자 포맷팅 (천 단위 구분자)
- ✅ 압축 숫자 (1K, 1.5M)
- ✅ 텍스트 자르기
- ✅ 마크다운 발췌문 추출
- ✅ 파일명 sanitize
- ✅ 파일 크기 포맷팅
- ✅ 슬러그 생성
- ✅ 태그 파싱/포맷팅
- ✅ 읽기 시간 계산

### collection.test.ts
- ✅ 배열 중복 제거 (unique, uniqueBy)
- ✅ 그룹화 (groupBy)
- ✅ 청크 분할 (chunk)
- ✅ 정렬 (sortBy)
- ✅ 평탄화 (flatten)
- ✅ 랜덤 샘플 (sample, sampleSize)
- ✅ 섞기 (shuffle)
- ✅ 빈 값 확인 (isEmpty)
- ✅ 객체 조작 (omit, pick)
- ✅ 깊은 복사 (cloneDeep)
- ✅ 깊은 병합 (mergeDeep)
- ✅ 중첩 값 접근 (get, set)
- ✅ Falsy 제거 (compact)
- ✅ 집합 연산 (intersection, difference)
- ✅ 범위 생성 (range)

## 📝 테스트 작성 가이드

### 기본 구조

```typescript
import { describe, it, expect } from 'vitest';
import { myFunction } from '../myUtility';

describe('myFunction', () => {
  it('should do something', () => {
    const result = myFunction(input);
    expect(result).toBe(expected);
  });

  it('should handle edge case', () => {
    expect(myFunction(edgeCase)).toBe(expectedEdgeCase);
  });
});
```

### 매처 (Matchers)

```typescript
// 동등성
expect(value).toBe(expected);
expect(value).toEqual(expected);

// 참/거짓
expect(value).toBeTruthy();
expect(value).toBeFalsy();

// 배열/객체
expect(array).toHaveLength(3);
expect(array).toContain(item);
expect(object).toHaveProperty('key');

// 예외
expect(() => fn()).toThrow();
```

## 🎯 테스트 커버리지 목표

- **전체**: 80% 이상
- **핵심 유틸리티**: 90% 이상
- **비즈니스 로직**: 95% 이상

## 🔧 설정 파일

### vitest.config.ts
```typescript
import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    globals: true,
    environment: 'happy-dom',
    setupFiles: ['./lib/test/setup.ts'],
    coverage: {
      provider: 'v8',
      reporter: ['text', 'json', 'html']
    }
  }
});
```

### lib/test/setup.ts
- jest-dom 매처 확장
- 전역 Mock 설정
- 테스트 환경 초기화

## 📊 테스트 실행 예제

```bash
# 특정 파일만 테스트
npm test validation.test.ts

# 패턴 매칭으로 테스트
npm test -- --grep "validation"

# 상세 모드로 실행
npm test -- --reporter=verbose

# 특정 브라우저 환경 테스트
npm test -- --browser.name=chrome
```

## 🐛 디버깅

### VS Code 디버거 설정

`.vscode/launch.json`:
```json
{
  "type": "node",
  "request": "launch",
  "name": "Debug Vitest Tests",
  "runtimeExecutable": "npm",
  "runtimeArgs": ["test", "--", "--run"],
  "console": "integratedTerminal"
}
```

### 콘솔 로그 확인
```typescript
it('should debug', () => {
  console.log('Debug info:', value);
  expect(value).toBe(expected);
});
```

## 📚 추가 리소스

- [Vitest 문서](https://vitest.dev/)
- [Testing Library 문서](https://testing-library.com/)
- [Jest-DOM 매처](https://github.com/testing-library/jest-dom)

## 🤝 기여 가이드

1. 새로운 유틸리티 함수를 추가할 때는 반드시 테스트를 작성하세요
2. 최소 3가지 케이스를 테스트하세요:
   - 정상 동작
   - 엣지 케이스
   - 에러 케이스
3. 테스트는 독립적이어야 하며 순서에 의존하면 안 됩니다
4. 의미 있는 테스트 설명을 작성하세요 (한국어 가능)

## ⚠️ 주의사항

- 테스트는 격리되어야 합니다 (다른 테스트에 영향 X)
- 외부 API나 DB에 의존하지 마세요
- 시간 관련 테스트는 고정된 날짜를 사용하세요
- 랜덤 값은 테스트하기 어려우므로 시드를 사용하세요

/**
 * Collection utilities tests
 */

import { describe, it, expect } from 'vitest';
import {
  unique,
  uniqueBy,
  groupBy,
  chunk,
  sortBy,
  flatten,
  sample,
  sampleSize,
  shuffle,
  isEmpty,
  omit,
  pick,
  cloneDeep,
  mergeDeep,
  get,
  set,
  compact,
  intersection,
  difference,
  range
} from '../collection';

describe('unique', () => {
  it('중복을 제거해야 함', () => {
    expect(unique([1, 2, 2, 3, 3, 3])).toEqual([1, 2, 3]);
    expect(unique(['a', 'b', 'a', 'c'])).toEqual(['a', 'b', 'c']);
  });

  it('빈 배열을 처리해야 함', () => {
    expect(unique([])).toEqual([]);
  });
});

describe('uniqueBy', () => {
  it('키로 중복을 제거해야 함', () => {
    const data = [
      { id: 1, name: 'A' },
      { id: 2, name: 'B' },
      { id: 1, name: 'C' }
    ];
    const result = uniqueBy(data, 'id');
    expect(result).toHaveLength(2);
    expect(result[0].id).toBe(1);
    expect(result[1].id).toBe(2);
  });
});

describe('groupBy', () => {
  it('키로 그룹화해야 함', () => {
    const data = [
      { category: 'A', value: 1 },
      { category: 'B', value: 2 },
      { category: 'A', value: 3 }
    ];
    const grouped = groupBy(data, 'category');
    expect(grouped['A']).toHaveLength(2);
    expect(grouped['B']).toHaveLength(1);
  });
});

describe('chunk', () => {
  it('배열을 청크로 나눠야 함', () => {
    const result = chunk([1, 2, 3, 4, 5], 2);
    expect(result).toEqual([[1, 2], [3, 4], [5]]);
  });

  it('빈 배열을 처리해야 함', () => {
    expect(chunk([], 2)).toEqual([]);
  });
});

describe('sortBy', () => {
  it('키로 오름차순 정렬해야 함', () => {
    const data = [
      { age: 30 },
      { age: 20 },
      { age: 25 }
    ];
    const sorted = sortBy(data, 'age', 'asc');
    expect(sorted[0].age).toBe(20);
    expect(sorted[2].age).toBe(30);
  });

  it('키로 내림차순 정렬해야 함', () => {
    const data = [
      { age: 30 },
      { age: 20 },
      { age: 25 }
    ];
    const sorted = sortBy(data, 'age', 'desc');
    expect(sorted[0].age).toBe(30);
    expect(sorted[2].age).toBe(20);
  });
});

describe('flatten', () => {
  it('중첩 배열을 평탄화해야 함', () => {
    expect(flatten([1, [2, [3, [4]]]])).toEqual([1, 2, 3, 4]);
  });

  it('이미 평탄화된 배열을 그대로 반환해야 함', () => {
    expect(flatten([1, 2, 3])).toEqual([1, 2, 3]);
  });
});

describe('sample', () => {
  it('배열에서 하나를 선택해야 함', () => {
    const arr = [1, 2, 3, 4, 5];
    const result = sample(arr);
    expect(arr).toContain(result);
  });

  it('빈 배열은 undefined를 반환해야 함', () => {
    expect(sample([])).toBeUndefined();
  });
});

describe('sampleSize', () => {
  it('배열에서 n개를 선택해야 함', () => {
    const arr = [1, 2, 3, 4, 5];
    const result = sampleSize(arr, 3);
    expect(result).toHaveLength(3);
    result.forEach(item => expect(arr).toContain(item));
  });

  it('배열보다 큰 크기를 요청하면 전체 배열을 반환해야 함', () => {
    const arr = [1, 2, 3];
    const result = sampleSize(arr, 10);
    expect(result).toHaveLength(3);
  });
});

describe('shuffle', () => {
  it('배열을 섞어야 함', () => {
    const arr = [1, 2, 3, 4, 5];
    const shuffled = shuffle(arr);
    expect(shuffled).toHaveLength(arr.length);
    shuffled.forEach(item => expect(arr).toContain(item));
  });

  it('원본 배열을 변경하지 않아야 함', () => {
    const arr = [1, 2, 3];
    const shuffled = shuffle(arr);
    expect(arr).toEqual([1, 2, 3]);
  });
});

describe('isEmpty', () => {
  it('빈 배열을 감지해야 함', () => {
    expect(isEmpty([])).toBe(true);
  });

  it('빈 객체를 감지해야 함', () => {
    expect(isEmpty({})).toBe(true);
  });

  it('빈 문자열을 감지해야 함', () => {
    expect(isEmpty('')).toBe(true);
  });

  it('null과 undefined를 감지해야 함', () => {
    expect(isEmpty(null)).toBe(true);
    expect(isEmpty(undefined)).toBe(true);
  });

  it('값이 있는 경우 false를 반환해야 함', () => {
    expect(isEmpty([1])).toBe(false);
    expect(isEmpty({ a: 1 })).toBe(false);
    expect(isEmpty('text')).toBe(false);
  });
});

describe('omit', () => {
  it('키를 제외해야 함', () => {
    const obj = { a: 1, b: 2, c: 3 };
    const result = omit(obj, ['b']);
    expect(result).toEqual({ a: 1, c: 3 });
  });
});

describe('pick', () => {
  it('키만 선택해야 함', () => {
    const obj = { a: 1, b: 2, c: 3 };
    const result = pick(obj, ['a', 'c']);
    expect(result).toEqual({ a: 1, c: 3 });
  });
});

describe('cloneDeep', () => {
  it('깊은 복사를 해야 함', () => {
    const obj = { a: 1, b: { c: 2 } };
    const cloned = cloneDeep(obj);
    cloned.b.c = 3;
    expect(obj.b.c).toBe(2);
  });

  it('배열을 깊은 복사해야 함', () => {
    const arr = [1, [2, 3]];
    const cloned = cloneDeep(arr);
    (cloned[1] as number[])[0] = 4;
    expect((arr[1] as number[])[0]).toBe(2);
  });

  it('Date 객체를 복사해야 함', () => {
    const date = new Date('2024-01-01');
    const cloned = cloneDeep(date);
    expect(cloned.getTime()).toBe(date.getTime());
  });
});

describe('mergeDeep', () => {
  it('객체를 깊게 병합해야 함', () => {
    const target = { a: 1, b: { c: 2 } };
    const source = { b: { d: 3 }, e: 4 };
    const result = mergeDeep(target, source);
    expect(result).toEqual({ a: 1, b: { c: 2, d: 3 }, e: 4 });
  });
});

describe('get', () => {
  it('중첩된 값을 가져와야 함', () => {
    const obj = { a: { b: { c: 1 } } };
    expect(get(obj, 'a.b.c')).toBe(1);
  });

  it('존재하지 않는 경로는 기본값을 반환해야 함', () => {
    const obj = { a: 1 };
    expect(get(obj, 'b.c', 'default')).toBe('default');
  });
});

describe('set', () => {
  it('중첩된 값을 설정해야 함', () => {
    const obj = { a: 1 };
    set(obj, 'b.c', 2);
    expect(obj).toEqual({ a: 1, b: { c: 2 } });
  });
});

describe('compact', () => {
  it('falsy 값을 제거해야 함', () => {
    expect(compact([1, 0, 2, false, 3, null, 4, undefined, 5, ''])).toEqual([1, 2, 3, 4, 5]);
  });
});

describe('intersection', () => {
  it('교집합을 반환해야 함', () => {
    expect(intersection([1, 2, 3], [2, 3, 4])).toEqual([2, 3]);
  });

  it('여러 배열의 교집합을 반환해야 함', () => {
    expect(intersection([1, 2, 3], [2, 3, 4], [2, 5])).toEqual([2]);
  });
});

describe('difference', () => {
  it('차집합을 반환해야 함', () => {
    expect(difference([1, 2, 3], [2, 3, 4])).toEqual([1]);
  });

  it('여러 배열의 차집합을 반환해야 함', () => {
    expect(difference([1, 2, 3, 4], [2], [3])).toEqual([1, 4]);
  });
});

describe('range', () => {
  it('범위를 생성해야 함', () => {
    expect(range(5)).toEqual([0, 1, 2, 3, 4]);
    expect(range(1, 5)).toEqual([1, 2, 3, 4]);
  });

  it('스텝을 지원해야 함', () => {
    expect(range(0, 10, 2)).toEqual([0, 2, 4, 6, 8]);
  });
});

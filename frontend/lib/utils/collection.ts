/**
 * Collection utilities for Viblog frontend
 * Array and object manipulation helpers
 */

/**
 * Remove duplicates from array
 */
export function unique<T>(array: T[]): T[] {
  return [...new Set(array)];
}

/**
 * Remove duplicates by key
 */
export function uniqueBy<T>(array: T[], key: keyof T): T[] {
  const seen = new Set();
  return array.filter(item => {
    const value = item[key];
    if (seen.has(value)) {
      return false;
    }
    seen.add(value);
    return true;
  });
}

/**
 * Group array by key
 */
export function groupBy<T>(array: T[], key: keyof T): Record<string, T[]> {
  return array.reduce((result, item) => {
    const groupKey = String(item[key]);
    if (!result[groupKey]) {
      result[groupKey] = [];
    }
    result[groupKey].push(item);
    return result;
  }, {} as Record<string, T[]>);
}

/**
 * Chunk array into smaller arrays
 */
export function chunk<T>(array: T[], size: number): T[][] {
  const chunks: T[][] = [];
  for (let i = 0; i < array.length; i += size) {
    chunks.push(array.slice(i, i + size));
  }
  return chunks;
}

/**
 * Sort array by key
 */
export function sortBy<T>(array: T[], key: keyof T, order: 'asc' | 'desc' = 'asc'): T[] {
  return [...array].sort((a, b) => {
    const aVal = a[key];
    const bVal = b[key];

    if (aVal < bVal) return order === 'asc' ? -1 : 1;
    if (aVal > bVal) return order === 'asc' ? 1 : -1;
    return 0;
  });
}

/**
 * Flatten nested array
 */
export function flatten<T>(array: any[]): T[] {
  return array.flat(Infinity) as T[];
}

/**
 * Get random element from array
 */
export function sample<T>(array: T[]): T | undefined {
  if (array.length === 0) return undefined;
  const randomIndex = Math.floor(Math.random() * array.length);
  return array[randomIndex];
}

/**
 * Get random elements from array
 */
export function sampleSize<T>(array: T[], size: number): T[] {
  const shuffled = [...array].sort(() => 0.5 - Math.random());
  return shuffled.slice(0, Math.min(size, array.length));
}

/**
 * Shuffle array
 */
export function shuffle<T>(array: T[]): T[] {
  const shuffled = [...array];
  for (let i = shuffled.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]];
  }
  return shuffled;
}

/**
 * Check if array is empty
 */
export function isEmpty(value: any): boolean {
  if (value == null) return true;
  if (Array.isArray(value) || typeof value === 'string') return value.length === 0;
  if (typeof value === 'object') return Object.keys(value).length === 0;
  return false;
}

/**
 * Omit keys from object
 */
export function omit<T extends object, K extends keyof T>(
  obj: T,
  keys: K[]
): Omit<T, K> {
  const result = { ...obj };
  keys.forEach(key => delete result[key]);
  return result;
}

/**
 * Pick keys from object
 */
export function pick<T extends object, K extends keyof T>(
  obj: T,
  keys: K[]
): Pick<T, K> {
  const result = {} as Pick<T, K>;
  keys.forEach(key => {
    if (key in obj) {
      result[key] = obj[key];
    }
  });
  return result;
}

/**
 * Deep clone object
 */
export function cloneDeep<T>(obj: T): T {
  if (obj === null || typeof obj !== 'object') {
    return obj;
  }

  if (obj instanceof Date) {
    return new Date(obj.getTime()) as any;
  }

  if (obj instanceof Array) {
    return obj.map(item => cloneDeep(item)) as any;
  }

  if (obj instanceof Object) {
    const clonedObj = {} as T;
    Object.keys(obj).forEach(key => {
      (clonedObj as any)[key] = cloneDeep((obj as any)[key]);
    });
    return clonedObj;
  }

  return obj;
}

/**
 * Deep merge objects
 */
export function mergeDeep<T extends object>(target: T, ...sources: Partial<T>[]): T {
  if (!sources.length) return target;

  const source = sources.shift();
  if (!source) return target;

  if (isObject(target) && isObject(source)) {
    Object.keys(source).forEach(key => {
      const sourceValue = (source as any)[key];
      const targetValue = (target as any)[key];

      if (isObject(sourceValue)) {
        if (!targetValue) {
          (target as any)[key] = {};
        }
        (target as any)[key] = mergeDeep(targetValue || {}, sourceValue);
      } else {
        (target as any)[key] = sourceValue;
      }
    });
  }

  return mergeDeep(target, ...sources);
}

/**
 * Check if value is object
 */
function isObject(item: any): boolean {
  return item && typeof item === 'object' && !Array.isArray(item);
}

/**
 * Get nested property value
 */
export function get<T>(obj: any, path: string, defaultValue?: T): T {
  const keys = path.split('.');
  let result = obj;

  for (const key of keys) {
    if (result == null) {
      return defaultValue as T;
    }
    result = result[key];
  }

  return result !== undefined ? result : (defaultValue as T);
}

/**
 * Set nested property value
 */
export function set<T extends object>(obj: T, path: string, value: any): T {
  const keys = path.split('.');
  const lastKey = keys.pop()!;
  let current: any = obj;

  for (const key of keys) {
    if (!current[key] || typeof current[key] !== 'object') {
      current[key] = {};
    }
    current = current[key];
  }

  current[lastKey] = value;
  return obj;
}

/**
 * Compact array (remove falsy values)
 */
export function compact<T>(array: (T | null | undefined | false | 0 | '')[]): T[] {
  return array.filter(Boolean) as T[];
}

/**
 * Intersection of arrays
 */
export function intersection<T>(...arrays: T[][]): T[] {
  if (arrays.length === 0) return [];
  if (arrays.length === 1) return arrays[0];

  return arrays[0].filter(item =>
    arrays.slice(1).every(arr => arr.includes(item))
  );
}

/**
 * Difference of arrays
 */
export function difference<T>(array: T[], ...others: T[][]): T[] {
  const othersFlat = flatten<T>(others);
  return array.filter(item => !othersFlat.includes(item));
}

/**
 * Range generator
 */
export function range(start: number, end?: number, step: number = 1): number[] {
  if (end === undefined) {
    end = start;
    start = 0;
  }

  const result: number[] = [];
  for (let i = start; i < end; i += step) {
    result.push(i);
  }

  return result;
}

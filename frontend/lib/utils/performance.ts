/**
 * Performance monitoring utilities for Viblog frontend
 * Web Vitals, performance metrics, optimization helpers
 */

export interface PerformanceMetrics {
  FCP?: number; // First Contentful Paint
  LCP?: number; // Largest Contentful Paint
  FID?: number; // First Input Delay
  CLS?: number; // Cumulative Layout Shift
  TTFB?: number; // Time to First Byte
}

/**
 * Measure performance timing
 */
export function measurePerformance(markName: string): () => number {
  const startMark = `${markName}-start`;
  const endMark = `${markName}-end`;
  const measureName = markName;

  performance.mark(startMark);

  return () => {
    performance.mark(endMark);
    performance.measure(measureName, startMark, endMark);

    const measure = performance.getEntriesByName(measureName)[0];
    const duration = measure.duration;

    // Cleanup
    performance.clearMarks(startMark);
    performance.clearMarks(endMark);
    performance.clearMeasures(measureName);

    return duration;
  };
}

/**
 * Get Web Vitals metrics
 */
export function getWebVitals(): PerformanceMetrics {
  const metrics: PerformanceMetrics = {};

  if (typeof window === 'undefined' || !window.performance) {
    return metrics;
  }

  // Time to First Byte
  const navigationTiming = performance.getEntriesByType('navigation')[0] as PerformanceNavigationTiming;
  if (navigationTiming) {
    metrics.TTFB = navigationTiming.responseStart - navigationTiming.requestStart;
  }

  // First Contentful Paint
  const paintEntries = performance.getEntriesByType('paint');
  const fcpEntry = paintEntries.find(entry => entry.name === 'first-contentful-paint');
  if (fcpEntry) {
    metrics.FCP = fcpEntry.startTime;
  }

  return metrics;
}

/**
 * Observe Largest Contentful Paint
 */
export function observeLCP(callback: (value: number) => void): () => void {
  if (typeof window === 'undefined' || !('PerformanceObserver' in window)) {
    return () => {};
  }

  try {
    const observer = new PerformanceObserver((list) => {
      const entries = list.getEntries();
      const lastEntry = entries[entries.length - 1] as any;
      callback(lastEntry.renderTime || lastEntry.loadTime);
    });

    observer.observe({ type: 'largest-contentful-paint', buffered: true });

    return () => observer.disconnect();
  } catch (error) {
    console.error('Failed to observe LCP:', error);
    return () => {};
  }
}

/**
 * Observe First Input Delay
 */
export function observeFID(callback: (value: number) => void): () => void {
  if (typeof window === 'undefined' || !('PerformanceObserver' in window)) {
    return () => {};
  }

  try {
    const observer = new PerformanceObserver((list) => {
      const entries = list.getEntries();
      entries.forEach((entry: any) => {
        callback(entry.processingStart - entry.startTime);
      });
    });

    observer.observe({ type: 'first-input', buffered: true });

    return () => observer.disconnect();
  } catch (error) {
    console.error('Failed to observe FID:', error);
    return () => {};
  }
}

/**
 * Observe Cumulative Layout Shift
 */
export function observeCLS(callback: (value: number) => void): () => void {
  if (typeof window === 'undefined' || !('PerformanceObserver' in window)) {
    return () => {};
  }

  try {
    let clsValue = 0;

    const observer = new PerformanceObserver((list) => {
      const entries = list.getEntries();
      entries.forEach((entry: any) => {
        if (!entry.hadRecentInput) {
          clsValue += entry.value;
          callback(clsValue);
        }
      });
    });

    observer.observe({ type: 'layout-shift', buffered: true });

    return () => observer.disconnect();
  } catch (error) {
    console.error('Failed to observe CLS:', error);
    return () => {};
  }
}

/**
 * Get page load time
 */
export function getPageLoadTime(): number | null {
  if (typeof window === 'undefined' || !window.performance) {
    return null;
  }

  const navigationTiming = performance.getEntriesByType('navigation')[0] as PerformanceNavigationTiming;
  if (navigationTiming) {
    return navigationTiming.loadEventEnd - navigationTiming.fetchStart;
  }

  return null;
}

/**
 * Get DOM ready time
 */
export function getDOMReadyTime(): number | null {
  if (typeof window === 'undefined' || !window.performance) {
    return null;
  }

  const navigationTiming = performance.getEntriesByType('navigation')[0] as PerformanceNavigationTiming;
  if (navigationTiming) {
    return navigationTiming.domContentLoadedEventEnd - navigationTiming.fetchStart;
  }

  return null;
}

/**
 * Get resource timing
 */
export function getResourceTiming(resourceUrl: string): PerformanceResourceTiming | null {
  if (typeof window === 'undefined' || !window.performance) {
    return null;
  }

  const resources = performance.getEntriesByType('resource') as PerformanceResourceTiming[];
  return resources.find(r => r.name.includes(resourceUrl)) || null;
}

/**
 * Get all resource timings
 */
export function getAllResourceTimings(): PerformanceResourceTiming[] {
  if (typeof window === 'undefined' || !window.performance) {
    return [];
  }

  return performance.getEntriesByType('resource') as PerformanceResourceTiming[];
}

/**
 * Get slow resources (loading time > threshold)
 */
export function getSlowResources(thresholdMs: number = 1000): Array<{
  name: string;
  duration: number;
  size: number;
}> {
  const resources = getAllResourceTimings();

  return resources
    .filter(r => r.duration > thresholdMs)
    .map(r => ({
      name: r.name,
      duration: r.duration,
      size: r.transferSize || 0
    }))
    .sort((a, b) => b.duration - a.duration);
}

/**
 * Get memory usage (if available)
 */
export function getMemoryUsage(): {
  usedJSHeapSize?: number;
  totalJSHeapSize?: number;
  jsHeapSizeLimit?: number;
} | null {
  if (typeof window === 'undefined') {
    return null;
  }

  const memory = (performance as any).memory;
  if (!memory) {
    return null;
  }

  return {
    usedJSHeapSize: memory.usedJSHeapSize,
    totalJSHeapSize: memory.totalJSHeapSize,
    jsHeapSizeLimit: memory.jsHeapSizeLimit
  };
}

/**
 * Check if device has low memory
 */
export function hasLowMemory(): boolean {
  if (typeof navigator === 'undefined') {
    return false;
  }

  const deviceMemory = (navigator as any).deviceMemory;
  return deviceMemory ? deviceMemory < 4 : false;
}

/**
 * Check network connection type
 */
export function getNetworkType(): string | null {
  if (typeof navigator === 'undefined') {
    return null;
  }

  const connection = (navigator as any).connection || (navigator as any).mozConnection || (navigator as any).webkitConnection;
  return connection ? connection.effectiveType : null;
}

/**
 * Check if slow network
 */
export function isSlowNetwork(): boolean {
  const networkType = getNetworkType();
  return networkType === 'slow-2g' || networkType === '2g';
}

/**
 * Prefetch resource
 */
export function prefetchResource(url: string): void {
  if (typeof document === 'undefined') {
    return;
  }

  const link = document.createElement('link');
  link.rel = 'prefetch';
  link.href = url;
  document.head.appendChild(link);
}

/**
 * Preload resource
 */
export function preloadResource(url: string, as: string = 'fetch'): void {
  if (typeof document === 'undefined') {
    return;
  }

  const link = document.createElement('link');
  link.rel = 'preload';
  link.href = url;
  link.as = as;
  document.head.appendChild(link);
}

/**
 * Report performance metrics
 */
export function reportPerformanceMetrics(): void {
  if (typeof window === 'undefined') {
    return;
  }

  // Wait for page load
  window.addEventListener('load', () => {
    setTimeout(() => {
      const metrics = getWebVitals();
      const pageLoadTime = getPageLoadTime();
      const domReadyTime = getDOMReadyTime();

      console.log('📊 Performance Metrics:', {
        ...metrics,
        pageLoadTime,
        domReadyTime
      });

      // Send to analytics or monitoring service
    }, 0);
  });
}

/**
 * Monitor long tasks
 */
export function monitorLongTasks(callback: (duration: number) => void): () => void {
  if (typeof window === 'undefined' || !('PerformanceObserver' in window)) {
    return () => {};
  }

  try {
    const observer = new PerformanceObserver((list) => {
      const entries = list.getEntries();
      entries.forEach((entry) => {
        callback(entry.duration);
      });
    });

    observer.observe({ type: 'longtask', buffered: true });

    return () => observer.disconnect();
  } catch (error) {
    console.error('Failed to monitor long tasks:', error);
    return () => {};
  }
}

/**
 * Create performance report
 */
export function createPerformanceReport(): {
  metrics: PerformanceMetrics;
  pageLoadTime: number | null;
  domReadyTime: number | null;
  slowResources: Array<{ name: string; duration: number; size: number }>;
  memoryUsage: ReturnType<typeof getMemoryUsage>;
  networkType: string | null;
} {
  return {
    metrics: getWebVitals(),
    pageLoadTime: getPageLoadTime(),
    domReadyTime: getDOMReadyTime(),
    slowResources: getSlowResources(1000),
    memoryUsage: getMemoryUsage(),
    networkType: getNetworkType()
  };
}

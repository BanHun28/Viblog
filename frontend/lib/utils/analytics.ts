/**
 * Analytics and tracking utilities for Viblog frontend
 * Page views, events, user interactions tracking
 */

export interface AnalyticsEvent {
  category: string;
  action: string;
  label?: string;
  value?: number;
}

export interface PageViewData {
  path: string;
  title: string;
  referrer?: string;
  timestamp: number;
}

/**
 * Track page view
 */
export function trackPageView(path: string, title: string): void {
  const data: PageViewData = {
    path,
    title,
    referrer: document.referrer,
    timestamp: Date.now()
  };

  // Log to console in development
  if (process.env.NODE_ENV === 'development') {
    console.log('📊 Page View:', data);
  }

  // Send to analytics service (placeholder)
  sendToAnalytics('pageview', data);
}

/**
 * Track custom event
 */
export function trackEvent(event: AnalyticsEvent): void {
  if (process.env.NODE_ENV === 'development') {
    console.log('📊 Event:', event);
  }

  sendToAnalytics('event', event);
}

/**
 * Track post view
 */
export function trackPostView(postId: string, postTitle: string): void {
  trackEvent({
    category: 'Post',
    action: 'View',
    label: postTitle,
    value: parseInt(postId, 10) || undefined
  });
}

/**
 * Track post like
 */
export function trackPostLike(postId: string, postTitle: string): void {
  trackEvent({
    category: 'Post',
    action: 'Like',
    label: postTitle,
    value: parseInt(postId, 10) || undefined
  });
}

/**
 * Track comment submit
 */
export function trackCommentSubmit(postId: string, commentType: 'member' | 'anonymous'): void {
  trackEvent({
    category: 'Comment',
    action: 'Submit',
    label: commentType,
    value: parseInt(postId, 10) || undefined
  });
}

/**
 * Track search
 */
export function trackSearch(query: string, resultCount: number): void {
  trackEvent({
    category: 'Search',
    action: 'Query',
    label: query,
    value: resultCount
  });
}

/**
 * Track user registration
 */
export function trackUserRegistration(method: 'email'): void {
  trackEvent({
    category: 'User',
    action: 'Register',
    label: method
  });
}

/**
 * Track user login
 */
export function trackUserLogin(method: 'email'): void {
  trackEvent({
    category: 'User',
    action: 'Login',
    label: method
  });
}

/**
 * Track user logout
 */
export function trackUserLogout(): void {
  trackEvent({
    category: 'User',
    action: 'Logout'
  });
}

/**
 * Track share action
 */
export function trackShare(platform: string, contentType: 'post', contentId: string): void {
  trackEvent({
    category: 'Share',
    action: platform,
    label: contentType,
    value: parseInt(contentId, 10) || undefined
  });
}

/**
 * Track bookmark action
 */
export function trackBookmark(postId: string, action: 'add' | 'remove'): void {
  trackEvent({
    category: 'Bookmark',
    action,
    value: parseInt(postId, 10) || undefined
  });
}

/**
 * Track navigation click
 */
export function trackNavigation(destination: string): void {
  trackEvent({
    category: 'Navigation',
    action: 'Click',
    label: destination
  });
}

/**
 * Track error
 */
export function trackError(error: Error, context?: string): void {
  trackEvent({
    category: 'Error',
    action: error.name,
    label: context ? `${context}: ${error.message}` : error.message
  });

  // Also log to console
  console.error('Tracked error:', error, context);
}

/**
 * Track performance timing
 */
export function trackTiming(
  category: string,
  variable: string,
  timeMs: number,
  label?: string
): void {
  trackEvent({
    category: `Timing/${category}`,
    action: variable,
    label,
    value: Math.round(timeMs)
  });
}

/**
 * Track scroll depth
 */
export function trackScrollDepth(depth: number): void {
  trackEvent({
    category: 'Engagement',
    action: 'Scroll',
    label: `${depth}%`,
    value: depth
  });
}

/**
 * Track time on page
 */
export function trackTimeOnPage(seconds: number): void {
  trackEvent({
    category: 'Engagement',
    action: 'TimeOnPage',
    value: Math.round(seconds)
  });
}

/**
 * Track outbound link click
 */
export function trackOutboundLink(url: string): void {
  trackEvent({
    category: 'Outbound',
    action: 'Click',
    label: url
  });
}

/**
 * Send data to analytics service
 */
function sendToAnalytics(type: string, data: any): void {
  // Placeholder for actual analytics implementation
  // Could integrate with Google Analytics, Mixpanel, etc.

  // Example: Google Analytics 4
  if (typeof window !== 'undefined' && (window as any).gtag) {
    if (type === 'pageview') {
      (window as any).gtag('config', 'GA_MEASUREMENT_ID', {
        page_path: data.path,
        page_title: data.title
      });
    } else if (type === 'event') {
      (window as any).gtag('event', data.action, {
        event_category: data.category,
        event_label: data.label,
        value: data.value
      });
    }
  }
}

/**
 * Initialize analytics
 */
export function initializeAnalytics(): void {
  if (typeof window === 'undefined') {
    return;
  }

  // Track initial page view
  trackPageView(window.location.pathname, document.title);

  // Track scroll depth
  let maxScrollDepth = 0;
  const checkScrollDepth = () => {
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    const scrollHeight = document.documentElement.scrollHeight - window.innerHeight;
    const scrollDepth = Math.round((scrollTop / scrollHeight) * 100);

    if (scrollDepth > maxScrollDepth && scrollDepth % 25 === 0) {
      maxScrollDepth = scrollDepth;
      trackScrollDepth(scrollDepth);
    }
  };

  window.addEventListener('scroll', checkScrollDepth, { passive: true });

  // Track time on page
  const startTime = Date.now();
  window.addEventListener('beforeunload', () => {
    const timeOnPage = (Date.now() - startTime) / 1000;
    trackTimeOnPage(timeOnPage);
  });
}

/**
 * Get session data
 */
export function getSessionData(): {
  sessionId: string;
  startTime: number;
  pageViews: number;
} {
  const sessionKey = 'analytics_session';
  const stored = localStorage.getItem(sessionKey);

  if (stored) {
    try {
      return JSON.parse(stored);
    } catch {
      // Invalid data, create new session
    }
  }

  // Create new session
  const session = {
    sessionId: generateSessionId(),
    startTime: Date.now(),
    pageViews: 0
  };

  localStorage.setItem(sessionKey, JSON.stringify(session));
  return session;
}

/**
 * Generate session ID
 */
function generateSessionId(): string {
  return `${Date.now()}_${Math.random().toString(36).substring(2, 9)}`;
}

/**
 * Increment page view count
 */
export function incrementPageViewCount(): void {
  const session = getSessionData();
  session.pageViews++;
  localStorage.setItem('analytics_session', JSON.stringify(session));
}

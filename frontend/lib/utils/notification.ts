/**
 * Notification utilities for Viblog frontend
 * Toast notifications, browser notifications
 */

export type NotificationLevel = 'success' | 'error' | 'warning' | 'info';

export interface ToastOptions {
  duration?: number;
  position?: 'top-right' | 'top-left' | 'bottom-right' | 'bottom-left' | 'top-center' | 'bottom-center';
  dismissible?: boolean;
  icon?: string;
}

/**
 * Check if browser notifications are supported
 */
export function areBrowserNotificationsSupported(): boolean {
  return 'Notification' in window;
}

/**
 * Get browser notification permission status
 */
export function getNotificationPermission(): NotificationPermission {
  if (!areBrowserNotificationsSupported()) {
    return 'denied';
  }
  return Notification.permission;
}

/**
 * Request browser notification permission
 */
export async function requestNotificationPermission(): Promise<NotificationPermission> {
  if (!areBrowserNotificationsSupported()) {
    return 'denied';
  }

  try {
    return await Notification.requestPermission();
  } catch (error) {
    console.error('Failed to request notification permission:', error);
    return 'denied';
  }
}

/**
 * Show browser notification
 */
export function showBrowserNotification(
  title: string,
  options?: NotificationOptions
): Notification | null {
  if (!areBrowserNotificationsSupported()) {
    console.warn('Browser notifications not supported');
    return null;
  }

  if (Notification.permission !== 'granted') {
    console.warn('Notification permission not granted');
    return null;
  }

  try {
    return new Notification(title, {
      icon: '/favicon.ico',
      badge: '/favicon.ico',
      ...options
    });
  } catch (error) {
    console.error('Failed to show notification:', error);
    return null;
  }
}

/**
 * Show comment notification
 */
export function showCommentNotification(
  authorName: string,
  postTitle: string,
  onClick?: () => void
): Notification | null {
  const notification = showBrowserNotification(
    '새로운 댓글',
    {
      body: `${authorName}님이 "${postTitle}"에 댓글을 남겼습니다`,
      tag: 'comment',
      requireInteraction: false
    }
  );

  if (notification && onClick) {
    notification.onclick = () => {
      onClick();
      notification.close();
    };
  }

  return notification;
}

/**
 * Show reply notification
 */
export function showReplyNotification(
  authorName: string,
  onClick?: () => void
): Notification | null {
  const notification = showBrowserNotification(
    '새로운 답글',
    {
      body: `${authorName}님이 답글을 남겼습니다`,
      tag: 'reply',
      requireInteraction: false
    }
  );

  if (notification && onClick) {
    notification.onclick = () => {
      onClick();
      notification.close();
    };
  }

  return notification;
}

/**
 * Show like notification
 */
export function showLikeNotification(
  authorName: string,
  postTitle: string,
  onClick?: () => void
): Notification | null {
  const notification = showBrowserNotification(
    '좋아요',
    {
      body: `${authorName}님이 "${postTitle}"을 좋아합니다`,
      tag: 'like',
      requireInteraction: false
    }
  );

  if (notification && onClick) {
    notification.onclick = () => {
      onClick();
      notification.close();
    };
  }

  return notification;
}

/**
 * Vibrate device (mobile)
 */
export function vibrate(pattern: number | number[] = 200): boolean {
  if (!('vibrate' in navigator)) {
    return false;
  }

  try {
    navigator.vibrate(pattern);
    return true;
  } catch {
    return false;
  }
}

/**
 * Play notification sound
 */
export function playNotificationSound(soundUrl: string = '/sounds/notification.mp3'): void {
  try {
    const audio = new Audio(soundUrl);
    audio.volume = 0.5;
    audio.play().catch(error => {
      console.error('Failed to play notification sound:', error);
    });
  } catch (error) {
    console.error('Failed to create audio:', error);
  }
}

/**
 * Create toast notification message
 */
export function createToastMessage(
  message: string,
  level: NotificationLevel = 'info'
): {
  message: string;
  level: NotificationLevel;
  timestamp: number;
} {
  return {
    message,
    level,
    timestamp: Date.now()
  };
}

/**
 * Format notification time
 */
export function formatNotificationTime(timestamp: number): string {
  const now = Date.now();
  const diff = now - timestamp;

  const seconds = Math.floor(diff / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  if (seconds < 60) {
    return '방금 전';
  } else if (minutes < 60) {
    return `${minutes}분 전`;
  } else if (hours < 24) {
    return `${hours}시간 전`;
  } else if (days < 7) {
    return `${days}일 전`;
  } else {
    return new Date(timestamp).toLocaleDateString('ko-KR');
  }
}

/**
 * Group notifications by date
 */
export function groupNotificationsByDate(
  notifications: Array<{ timestamp: number; [key: string]: any }>
): Record<string, Array<{ timestamp: number; [key: string]: any }>> {
  const groups: Record<string, Array<{ timestamp: number; [key: string]: any }>> = {};

  notifications.forEach(notification => {
    const date = new Date(notification.timestamp);
    const today = new Date();
    const yesterday = new Date(today);
    yesterday.setDate(yesterday.getDate() - 1);

    let groupKey: string;

    if (date.toDateString() === today.toDateString()) {
      groupKey = '오늘';
    } else if (date.toDateString() === yesterday.toDateString()) {
      groupKey = '어제';
    } else {
      groupKey = date.toLocaleDateString('ko-KR');
    }

    if (!groups[groupKey]) {
      groups[groupKey] = [];
    }

    groups[groupKey].push(notification);
  });

  return groups;
}

/**
 * Check if notification is recent (within 24 hours)
 */
export function isRecentNotification(timestamp: number): boolean {
  const now = Date.now();
  const diff = now - timestamp;
  const hours = diff / (1000 * 60 * 60);

  return hours < 24;
}

/**
 * Mark notification as read
 */
export function markNotificationAsRead(notificationId: string): void {
  // This would typically interact with a store or API
  // Placeholder for implementation
  console.log(`Marking notification ${notificationId} as read`);
}

/**
 * Get unread notification count
 */
export function getUnreadCount(
  notifications: Array<{ read: boolean }>
): number {
  return notifications.filter(n => !n.read).length;
}

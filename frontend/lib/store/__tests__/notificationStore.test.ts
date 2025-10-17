import { describe, it, expect, beforeEach } from 'vitest'
import { useNotificationStore } from '../notificationStore'
import { Notification } from '@/types/notification'

describe('notificationStore', () => {
  const mockNotification: Notification = {
    id: '1',
    userId: 'user1',
    type: 'comment_reply',
    title: 'New Reply',
    content: 'Someone replied to your comment',
    relatedPostId: 'post1',
    relatedCommentId: 'comment1',
    isRead: false,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-01',
  }

  const mockReadNotification: Notification = {
    ...mockNotification,
    id: '2',
    isRead: true,
  }

  beforeEach(() => {
    useNotificationStore.setState({
      notifications: [],
      unreadCount: 0,
      preferences: {
        commentReply: true,
        postLike: true,
        commentLike: true,
      },
      isLoading: false,
      error: null,
    })
  })

  describe('Initial State', () => {
    it('should have correct initial state', () => {
      const state = useNotificationStore.getState()

      expect(state.notifications).toEqual([])
      expect(state.unreadCount).toBe(0)
      expect(state.preferences).toEqual({
        commentReply: true,
        postLike: true,
        commentLike: true,
      })
      expect(state.isLoading).toBe(false)
      expect(state.error).toBeNull()
    })
  })

  describe('Notifications Management', () => {
    it('should set notifications', () => {
      const notifications = [mockNotification, mockReadNotification]

      useNotificationStore.getState().setNotifications(notifications)

      const state = useNotificationStore.getState()
      expect(state.notifications).toEqual(notifications)
    })

    it('should calculate unread count when setting notifications', () => {
      const notifications = [
        mockNotification,
        mockReadNotification,
        { ...mockNotification, id: '3', isRead: false },
      ]

      useNotificationStore.getState().setNotifications(notifications)

      const state = useNotificationStore.getState()
      expect(state.unreadCount).toBe(2)
    })

    it('should add notification to beginning', () => {
      useNotificationStore.getState().setNotifications([mockReadNotification])

      useNotificationStore.getState().addNotification(mockNotification)

      const state = useNotificationStore.getState()
      expect(state.notifications).toHaveLength(2)
      expect(state.notifications[0]).toEqual(mockNotification)
    })

    it('should increment unread count when adding unread notification', () => {
      useNotificationStore.getState().addNotification(mockNotification)

      const state = useNotificationStore.getState()
      expect(state.unreadCount).toBe(1)
    })

    it('should not increment unread count when adding read notification', () => {
      useNotificationStore.getState().addNotification(mockReadNotification)

      const state = useNotificationStore.getState()
      expect(state.unreadCount).toBe(0)
    })

    it('should delete notification', () => {
      useNotificationStore.getState().setNotifications([mockNotification, mockReadNotification])

      useNotificationStore.getState().deleteNotification('1')

      const state = useNotificationStore.getState()
      expect(state.notifications).toHaveLength(1)
      expect(state.notifications[0].id).toBe('2')
    })

    it('should decrement unread count when deleting unread notification', () => {
      useNotificationStore.getState().setNotifications([mockNotification])

      useNotificationStore.getState().deleteNotification('1')

      const state = useNotificationStore.getState()
      expect(state.unreadCount).toBe(0)
    })

    it('should not decrement unread count when deleting read notification', () => {
      useNotificationStore.getState().setNotifications([mockNotification, mockReadNotification])

      useNotificationStore.getState().deleteNotification('2')

      const state = useNotificationStore.getState()
      expect(state.unreadCount).toBe(1)
    })
  })

  describe('Mark as Read', () => {
    it('should mark notification as read', () => {
      useNotificationStore.getState().setNotifications([mockNotification])

      useNotificationStore.getState().markAsRead('1')

      const state = useNotificationStore.getState()
      expect(state.notifications[0].isRead).toBe(true)
      expect(state.unreadCount).toBe(0)
    })

    it('should not change unread count if already read', () => {
      useNotificationStore.getState().setNotifications([mockReadNotification])

      useNotificationStore.getState().markAsRead('2')

      const state = useNotificationStore.getState()
      expect(state.unreadCount).toBe(0)
    })

    it('should mark all notifications as read', () => {
      const notifications = [
        mockNotification,
        { ...mockNotification, id: '2', isRead: false },
        { ...mockNotification, id: '3', isRead: false },
      ]

      useNotificationStore.getState().setNotifications(notifications)
      useNotificationStore.getState().markAllAsRead()

      const state = useNotificationStore.getState()
      expect(state.notifications.every((n) => n.isRead)).toBe(true)
      expect(state.unreadCount).toBe(0)
    })

    it('should handle markAllAsRead when no unread notifications', () => {
      useNotificationStore.getState().setNotifications([mockReadNotification])

      useNotificationStore.getState().markAllAsRead()

      const state = useNotificationStore.getState()
      expect(state.unreadCount).toBe(0)
    })
  })

  describe('Preferences Management', () => {
    it('should set preferences', () => {
      const newPreferences = {
        commentReply: false,
        postLike: false,
        commentLike: true,
      }

      useNotificationStore.getState().setPreferences(newPreferences)

      const state = useNotificationStore.getState()
      expect(state.preferences).toEqual(newPreferences)
    })

    it('should update single preference', () => {
      useNotificationStore.getState().updatePreference('commentReply', false)

      const state = useNotificationStore.getState()
      expect(state.preferences.commentReply).toBe(false)
      expect(state.preferences.postLike).toBe(true)
      expect(state.preferences.commentLike).toBe(true)
    })

    it('should update multiple preferences independently', () => {
      useNotificationStore.getState().updatePreference('commentReply', false)
      useNotificationStore.getState().updatePreference('postLike', false)

      const state = useNotificationStore.getState()
      expect(state.preferences.commentReply).toBe(false)
      expect(state.preferences.postLike).toBe(false)
      expect(state.preferences.commentLike).toBe(true)
    })
  })

  describe('Loading and Error States', () => {
    it('should set loading state', () => {
      useNotificationStore.getState().setLoading(true)
      expect(useNotificationStore.getState().isLoading).toBe(true)

      useNotificationStore.getState().setLoading(false)
      expect(useNotificationStore.getState().isLoading).toBe(false)
    })

    it('should set error state', () => {
      useNotificationStore.getState().setError('Test error')
      expect(useNotificationStore.getState().error).toBe('Test error')

      useNotificationStore.getState().setError(null)
      expect(useNotificationStore.getState().error).toBeNull()
    })
  })

  describe('Notification Types', () => {
    it('should handle comment_reply notification', () => {
      const notification: Notification = {
        ...mockNotification,
        type: 'comment_reply',
      }

      useNotificationStore.getState().addNotification(notification)

      const state = useNotificationStore.getState()
      expect(state.notifications[0].type).toBe('comment_reply')
    })

    it('should handle post_like notification', () => {
      const notification: Notification = {
        ...mockNotification,
        type: 'post_like',
      }

      useNotificationStore.getState().addNotification(notification)

      const state = useNotificationStore.getState()
      expect(state.notifications[0].type).toBe('post_like')
    })

    it('should handle comment_like notification', () => {
      const notification: Notification = {
        ...mockNotification,
        type: 'comment_like',
      }

      useNotificationStore.getState().addNotification(notification)

      const state = useNotificationStore.getState()
      expect(state.notifications[0].type).toBe('comment_like')
    })
  })

  describe('Complex Scenarios', () => {
    it('should handle multiple operations correctly', () => {
      // Add multiple notifications
      useNotificationStore.getState().addNotification(mockNotification)
      useNotificationStore.getState().addNotification({ ...mockNotification, id: '2' })
      useNotificationStore.getState().addNotification(mockReadNotification)

      expect(useNotificationStore.getState().unreadCount).toBe(2)

      // Mark one as read
      useNotificationStore.getState().markAsRead('1')
      expect(useNotificationStore.getState().unreadCount).toBe(1)

      // Delete one unread
      useNotificationStore.getState().deleteNotification('2')
      expect(useNotificationStore.getState().unreadCount).toBe(0)
    })

    it('should maintain correct state after multiple preference updates', () => {
      useNotificationStore.getState().updatePreference('commentReply', false)
      useNotificationStore.getState().updatePreference('commentReply', true)
      useNotificationStore.getState().updatePreference('postLike', false)

      const state = useNotificationStore.getState()
      expect(state.preferences.commentReply).toBe(true)
      expect(state.preferences.postLike).toBe(false)
    })
  })
})

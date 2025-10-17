import { create } from 'zustand'
import { Notification, NotificationPreferences } from '@/types/notification'

interface NotificationState {
  notifications: Notification[]
  unreadCount: number
  preferences: NotificationPreferences
  isLoading: boolean
  error: string | null
}

interface NotificationActions {
  setNotifications: (notifications: Notification[]) => void
  addNotification: (notification: Notification) => void
  markAsRead: (id: string) => void
  markAllAsRead: () => void
  deleteNotification: (id: string) => void
  setPreferences: (preferences: NotificationPreferences) => void
  updatePreference: (key: keyof NotificationPreferences, value: boolean) => void
  setLoading: (loading: boolean) => void
  setError: (error: string | null) => void
}

type NotificationStore = NotificationState & NotificationActions

export const useNotificationStore = create<NotificationStore>((set) => ({
  // State
  notifications: [],
  unreadCount: 0,
  preferences: {
    commentReply: true,
    postLike: true,
    commentLike: true,
  },
  isLoading: false,
  error: null,

  // Actions
  setNotifications: (notifications) =>
    set({
      notifications,
      unreadCount: notifications.filter((n) => !n.isRead).length,
    }),

  addNotification: (notification) =>
    set((state) => ({
      notifications: [notification, ...state.notifications],
      unreadCount: notification.isRead ? state.unreadCount : state.unreadCount + 1,
    })),

  markAsRead: (id) =>
    set((state) => ({
      notifications: state.notifications.map((n) =>
        n.id === id ? { ...n, isRead: true } : n
      ),
      unreadCount: state.notifications.find((n) => n.id === id && !n.isRead)
        ? state.unreadCount - 1
        : state.unreadCount,
    })),

  markAllAsRead: () =>
    set((state) => ({
      notifications: state.notifications.map((n) => ({ ...n, isRead: true })),
      unreadCount: 0,
    })),

  deleteNotification: (id) =>
    set((state) => ({
      notifications: state.notifications.filter((n) => n.id !== id),
      unreadCount: state.notifications.find((n) => n.id === id && !n.isRead)
        ? state.unreadCount - 1
        : state.unreadCount,
    })),

  setPreferences: (preferences) => set({ preferences }),

  updatePreference: (key, value) =>
    set((state) => ({
      preferences: { ...state.preferences, [key]: value },
    })),

  setLoading: (loading) => set({ isLoading: loading }),

  setError: (error) => set({ error }),
}))

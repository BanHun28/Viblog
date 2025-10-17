export type NotificationType = 'comment_reply' | 'post_like' | 'comment_like'

export interface Notification {
  id: string
  userId: string
  type: NotificationType
  title: string
  content: string
  relatedPostId?: string
  relatedCommentId?: string
  isRead: boolean
  createdAt: string
  updatedAt: string
}

export interface NotificationPreferences {
  commentReply: boolean
  postLike: boolean
  commentLike: boolean
}

import { create } from 'zustand'
import { Comment } from '@/types/comment'

interface CommentState {
  comments: Comment[]
  isLoading: boolean
  error: string | null
  replyingTo: string | null // comment ID being replied to
}

interface CommentActions {
  setComments: (comments: Comment[]) => void
  addComment: (comment: Comment) => void
  updateComment: (id: string, updates: Partial<Comment>) => void
  deleteComment: (id: string) => void
  setLoading: (loading: boolean) => void
  setError: (error: string | null) => void
  setReplyingTo: (commentId: string | null) => void
  toggleLike: (commentId: string, isLiked: boolean) => void
  // Helper for nested replies
  addReply: (parentId: string, reply: Comment) => void
}

type CommentStore = CommentState & CommentActions

export const useCommentStore = create<CommentStore>((set) => ({
  // State
  comments: [],
  isLoading: false,
  error: null,
  replyingTo: null,

  // Actions
  setComments: (comments) => set({ comments }),

  addComment: (comment) =>
    set((state) => ({
      comments: comment.parentId
        ? state.comments // If it's a reply, it will be added via addReply
        : [comment, ...state.comments],
    })),

  updateComment: (id, updates) =>
    set((state) => ({
      comments: state.comments.map((comment) =>
        comment.id === id
          ? { ...comment, ...updates }
          : {
              ...comment,
              replies: comment.replies?.map((reply) =>
                reply.id === id ? { ...reply, ...updates } : reply
              ),
            }
      ),
    })),

  deleteComment: (id) =>
    set((state) => ({
      comments: state.comments
        .filter((comment) => comment.id !== id)
        .map((comment) => ({
          ...comment,
          replies: comment.replies?.filter((reply) => reply.id !== id),
        })),
    })),

  setLoading: (loading) => set({ isLoading: loading }),

  setError: (error) => set({ error }),

  setReplyingTo: (commentId) => set({ replyingTo: commentId }),

  toggleLike: (commentId, isLiked) =>
    set((state) => ({
      comments: state.comments.map((comment) =>
        comment.id === commentId
          ? { ...comment, likeCount: comment.likeCount + (isLiked ? 1 : -1) }
          : {
              ...comment,
              replies: comment.replies?.map((reply) =>
                reply.id === commentId
                  ? { ...reply, likeCount: reply.likeCount + (isLiked ? 1 : -1) }
                  : reply
              ),
            }
      ),
    })),

  addReply: (parentId, reply) =>
    set((state) => ({
      comments: state.comments.map((comment) =>
        comment.id === parentId
          ? {
              ...comment,
              replies: [...(comment.replies || []), reply],
              replyCount: comment.replyCount + 1,
            }
          : comment
      ),
    })),
}))

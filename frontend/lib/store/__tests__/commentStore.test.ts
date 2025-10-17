import { describe, it, expect, beforeEach } from 'vitest'
import { useCommentStore } from '../commentStore'
import { Comment } from '@/types/comment'

describe('commentStore', () => {
  const mockComment: Comment = {
    id: '1',
    postId: 'post1',
    content: 'Test comment',
    isAnonymous: false,
    authorId: 'user1',
    author: {
      id: 'user1',
      nickname: 'testuser',
    },
    likeCount: 0,
    replyCount: 0,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-01',
  }

  const mockAnonymousComment: Comment = {
    id: '2',
    postId: 'post1',
    content: 'Anonymous comment',
    isAnonymous: true,
    anonymousNickname: 'Anonymous',
    likeCount: 0,
    replyCount: 0,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-01',
  }

  beforeEach(() => {
    useCommentStore.setState({
      comments: [],
      isLoading: false,
      error: null,
      replyingTo: null,
    })
  })

  describe('Initial State', () => {
    it('should have correct initial state', () => {
      const state = useCommentStore.getState()

      expect(state.comments).toEqual([])
      expect(state.isLoading).toBe(false)
      expect(state.error).toBeNull()
      expect(state.replyingTo).toBeNull()
    })
  })

  describe('Comments Management', () => {
    it('should set comments', () => {
      const comments = [mockComment, mockAnonymousComment]

      useCommentStore.getState().setComments(comments)
      expect(useCommentStore.getState().comments).toEqual(comments)
    })

    it('should add top-level comment', () => {
      useCommentStore.getState().addComment(mockComment)

      const state = useCommentStore.getState()
      expect(state.comments).toHaveLength(1)
      expect(state.comments[0]).toEqual(mockComment)
    })

    it('should add comment to beginning of list', () => {
      const existingComment = { ...mockComment, id: '2' }
      useCommentStore.getState().setComments([existingComment])

      useCommentStore.getState().addComment(mockComment)

      const state = useCommentStore.getState()
      expect(state.comments[0]).toEqual(mockComment)
      expect(state.comments[1]).toEqual(existingComment)
    })

    it('should not add reply via addComment', () => {
      const reply: Comment = {
        ...mockComment,
        id: '3',
        parentId: '1',
      }

      useCommentStore.getState().setComments([mockComment])
      useCommentStore.getState().addComment(reply)

      // Reply with parentId should not be added to top-level
      const state = useCommentStore.getState()
      expect(state.comments).toHaveLength(1)
    })

    it('should update comment', () => {
      useCommentStore.getState().setComments([mockComment])

      useCommentStore.getState().updateComment('1', { content: 'Updated content' })

      const state = useCommentStore.getState()
      expect(state.comments[0].content).toBe('Updated content')
    })

    it('should update reply within comment', () => {
      const reply: Comment = {
        ...mockComment,
        id: '2',
        parentId: '1',
      }
      const commentWithReply: Comment = {
        ...mockComment,
        replies: [reply],
      }

      useCommentStore.getState().setComments([commentWithReply])
      useCommentStore.getState().updateComment('2', { content: 'Updated reply' })

      const state = useCommentStore.getState()
      expect(state.comments[0].replies?.[0].content).toBe('Updated reply')
    })

    it('should delete comment', () => {
      useCommentStore.getState().setComments([mockComment, mockAnonymousComment])

      useCommentStore.getState().deleteComment('1')

      const state = useCommentStore.getState()
      expect(state.comments).toHaveLength(1)
      expect(state.comments[0].id).toBe('2')
    })

    it('should delete reply', () => {
      const reply: Comment = {
        ...mockComment,
        id: '2',
        parentId: '1',
      }
      const commentWithReply: Comment = {
        ...mockComment,
        replies: [reply],
      }

      useCommentStore.getState().setComments([commentWithReply])
      useCommentStore.getState().deleteComment('2')

      const state = useCommentStore.getState()
      expect(state.comments[0].replies).toHaveLength(0)
    })
  })

  describe('Loading and Error States', () => {
    it('should set loading state', () => {
      useCommentStore.getState().setLoading(true)
      expect(useCommentStore.getState().isLoading).toBe(true)

      useCommentStore.getState().setLoading(false)
      expect(useCommentStore.getState().isLoading).toBe(false)
    })

    it('should set error state', () => {
      useCommentStore.getState().setError('Test error')
      expect(useCommentStore.getState().error).toBe('Test error')

      useCommentStore.getState().setError(null)
      expect(useCommentStore.getState().error).toBeNull()
    })
  })

  describe('Reply Management', () => {
    it('should set replyingTo', () => {
      useCommentStore.getState().setReplyingTo('1')
      expect(useCommentStore.getState().replyingTo).toBe('1')
    })

    it('should clear replyingTo', () => {
      useCommentStore.getState().setReplyingTo('1')
      useCommentStore.getState().setReplyingTo(null)
      expect(useCommentStore.getState().replyingTo).toBeNull()
    })

    it('should add reply to parent comment', () => {
      const reply: Comment = {
        ...mockComment,
        id: '2',
        parentId: '1',
        content: 'Reply content',
      }

      useCommentStore.getState().setComments([mockComment])
      useCommentStore.getState().addReply('1', reply)

      const state = useCommentStore.getState()
      expect(state.comments[0].replies).toHaveLength(1)
      expect(state.comments[0].replies?.[0]).toEqual(reply)
      expect(state.comments[0].replyCount).toBe(1)
    })

    it('should add multiple replies', () => {
      const reply1: Comment = {
        ...mockComment,
        id: '2',
        parentId: '1',
        content: 'Reply 1',
      }
      const reply2: Comment = {
        ...mockComment,
        id: '3',
        parentId: '1',
        content: 'Reply 2',
      }

      useCommentStore.getState().setComments([mockComment])
      useCommentStore.getState().addReply('1', reply1)
      useCommentStore.getState().addReply('1', reply2)

      const state = useCommentStore.getState()
      expect(state.comments[0].replies).toHaveLength(2)
      expect(state.comments[0].replyCount).toBe(2)
    })

    it('should not add reply to non-existent comment', () => {
      const reply: Comment = {
        ...mockComment,
        id: '2',
        parentId: '999',
      }

      useCommentStore.getState().setComments([mockComment])
      useCommentStore.getState().addReply('999', reply)

      const state = useCommentStore.getState()
      expect(state.comments[0].replies).toBeUndefined()
    })
  })

  describe('Like Functionality', () => {
    it('should toggle like on comment', () => {
      useCommentStore.getState().setComments([mockComment])

      useCommentStore.getState().toggleLike('1', true)

      const state = useCommentStore.getState()
      expect(state.comments[0].likeCount).toBe(1)
    })

    it('should decrement like count when unliking', () => {
      const commentWithLike = { ...mockComment, likeCount: 1 }
      useCommentStore.getState().setComments([commentWithLike])

      useCommentStore.getState().toggleLike('1', false)

      const state = useCommentStore.getState()
      expect(state.comments[0].likeCount).toBe(0)
    })

    it('should toggle like on reply', () => {
      const reply: Comment = {
        ...mockComment,
        id: '2',
        parentId: '1',
      }
      const commentWithReply: Comment = {
        ...mockComment,
        replies: [reply],
      }

      useCommentStore.getState().setComments([commentWithReply])
      useCommentStore.getState().toggleLike('2', true)

      const state = useCommentStore.getState()
      expect(state.comments[0].replies?.[0].likeCount).toBe(1)
    })

    it('should not affect other comments', () => {
      useCommentStore.getState().setComments([mockComment, mockAnonymousComment])

      useCommentStore.getState().toggleLike('1', true)

      const state = useCommentStore.getState()
      expect(state.comments[0].likeCount).toBe(1)
      expect(state.comments[1].likeCount).toBe(0)
    })
  })

  describe('Anonymous Comments', () => {
    it('should handle anonymous comments', () => {
      useCommentStore.getState().setComments([mockAnonymousComment])

      const state = useCommentStore.getState()
      expect(state.comments[0].isAnonymous).toBe(true)
      expect(state.comments[0].anonymousNickname).toBe('Anonymous')
      expect(state.comments[0].authorId).toBeUndefined()
    })

    it('should handle mixed anonymous and user comments', () => {
      useCommentStore.getState().setComments([mockComment, mockAnonymousComment])

      const state = useCommentStore.getState()
      expect(state.comments).toHaveLength(2)
      expect(state.comments[0].isAnonymous).toBe(false)
      expect(state.comments[1].isAnonymous).toBe(true)
    })
  })
})

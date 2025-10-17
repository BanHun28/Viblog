import { describe, it, expect, beforeEach } from 'vitest'
import { usePostStore } from '../postStore'
import { Post, Category, Tag } from '@/types/post'

describe('postStore', () => {
  const mockPost: Post = {
    id: '1',
    title: 'Test Post',
    content: 'Test Content',
    status: 'published',
    authorId: 'user1',
    viewCount: 0,
    likeCount: 0,
    commentCount: 0,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-01',
  }

  const mockCategory: Category = {
    id: '1',
    name: 'Tech',
    slug: 'tech',
    postCount: 0,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-01',
  }

  const mockTag: Tag = {
    id: '1',
    name: 'JavaScript',
    slug: 'javascript',
    postCount: 0,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-01',
  }

  beforeEach(() => {
    usePostStore.setState({
      posts: [],
      currentPost: null,
      categories: [],
      tags: [],
      isLoading: false,
      error: null,
      hasNext: false,
      nextCursor: null,
    })
  })

  describe('Initial State', () => {
    it('should have correct initial state', () => {
      const state = usePostStore.getState()

      expect(state.posts).toEqual([])
      expect(state.currentPost).toBeNull()
      expect(state.categories).toEqual([])
      expect(state.tags).toEqual([])
      expect(state.isLoading).toBe(false)
      expect(state.error).toBeNull()
      expect(state.hasNext).toBe(false)
      expect(state.nextCursor).toBeNull()
    })
  })

  describe('Posts Management', () => {
    it('should set posts', () => {
      const posts = [mockPost, { ...mockPost, id: '2', title: 'Post 2' }]

      usePostStore.getState().setPosts(posts)
      expect(usePostStore.getState().posts).toEqual(posts)
    })

    it('should add post to beginning of list', () => {
      const existingPost = { ...mockPost, id: '2' }
      usePostStore.getState().setPosts([existingPost])

      const newPost = { ...mockPost, id: '1' }
      usePostStore.getState().addPost(newPost)

      const state = usePostStore.getState()
      expect(state.posts).toHaveLength(2)
      expect(state.posts[0]).toEqual(newPost)
      expect(state.posts[1]).toEqual(existingPost)
    })

    it('should update post', () => {
      usePostStore.getState().setPosts([mockPost])

      usePostStore.getState().updatePost('1', { title: 'Updated Title' })

      const state = usePostStore.getState()
      expect(state.posts[0].title).toBe('Updated Title')
      expect(state.posts[0].content).toBe('Test Content')
    })

    it('should not update non-existent post', () => {
      usePostStore.getState().setPosts([mockPost])

      usePostStore.getState().updatePost('999', { title: 'Updated' })

      const state = usePostStore.getState()
      expect(state.posts[0].title).toBe('Test Post')
    })

    it('should delete post', () => {
      usePostStore.getState().setPosts([mockPost, { ...mockPost, id: '2' }])

      usePostStore.getState().deletePost('1')

      const state = usePostStore.getState()
      expect(state.posts).toHaveLength(1)
      expect(state.posts[0].id).toBe('2')
    })

    it('should clear posts', () => {
      usePostStore.getState().setPosts([mockPost])
      usePostStore.getState().setPagination(true, 'cursor123')

      usePostStore.getState().clearPosts()

      const state = usePostStore.getState()
      expect(state.posts).toEqual([])
      expect(state.hasNext).toBe(false)
      expect(state.nextCursor).toBeNull()
    })
  })

  describe('Current Post Management', () => {
    it('should set current post', () => {
      usePostStore.getState().setCurrentPost(mockPost)
      expect(usePostStore.getState().currentPost).toEqual(mockPost)
    })

    it('should clear current post', () => {
      usePostStore.getState().setCurrentPost(mockPost)
      usePostStore.getState().setCurrentPost(null)
      expect(usePostStore.getState().currentPost).toBeNull()
    })

    it('should update current post when updating posts', () => {
      usePostStore.getState().setPosts([mockPost])
      usePostStore.getState().setCurrentPost(mockPost)

      usePostStore.getState().updatePost('1', { title: 'Updated' })

      const state = usePostStore.getState()
      expect(state.currentPost?.title).toBe('Updated')
    })

    it('should clear current post when deleting it', () => {
      usePostStore.getState().setPosts([mockPost])
      usePostStore.getState().setCurrentPost(mockPost)

      usePostStore.getState().deletePost('1')

      expect(usePostStore.getState().currentPost).toBeNull()
    })

    it('should not clear current post when deleting different post', () => {
      const post2 = { ...mockPost, id: '2' }
      usePostStore.getState().setPosts([mockPost, post2])
      usePostStore.getState().setCurrentPost(mockPost)

      usePostStore.getState().deletePost('2')

      expect(usePostStore.getState().currentPost).toEqual(mockPost)
    })
  })

  describe('Categories Management', () => {
    it('should set categories', () => {
      const categories = [mockCategory, { ...mockCategory, id: '2', name: 'Design' }]

      usePostStore.getState().setCategories(categories)
      expect(usePostStore.getState().categories).toEqual(categories)
    })
  })

  describe('Tags Management', () => {
    it('should set tags', () => {
      const tags = [mockTag, { ...mockTag, id: '2', name: 'React' }]

      usePostStore.getState().setTags(tags)
      expect(usePostStore.getState().tags).toEqual(tags)
    })
  })

  describe('Loading and Error States', () => {
    it('should set loading state', () => {
      usePostStore.getState().setLoading(true)
      expect(usePostStore.getState().isLoading).toBe(true)

      usePostStore.getState().setLoading(false)
      expect(usePostStore.getState().isLoading).toBe(false)
    })

    it('should set error state', () => {
      usePostStore.getState().setError('Test error')
      expect(usePostStore.getState().error).toBe('Test error')

      usePostStore.getState().setError(null)
      expect(usePostStore.getState().error).toBeNull()
    })
  })

  describe('Pagination', () => {
    it('should set pagination data', () => {
      usePostStore.getState().setPagination(true, 'cursor123')

      const state = usePostStore.getState()
      expect(state.hasNext).toBe(true)
      expect(state.nextCursor).toBe('cursor123')
    })

    it('should handle last page', () => {
      usePostStore.getState().setPagination(false, null)

      const state = usePostStore.getState()
      expect(state.hasNext).toBe(false)
      expect(state.nextCursor).toBeNull()
    })
  })

  describe('Like Functionality', () => {
    it('should increment like count when liking', () => {
      usePostStore.getState().setPosts([mockPost])

      usePostStore.getState().toggleLike('1', true)

      const state = usePostStore.getState()
      expect(state.posts[0].likeCount).toBe(1)
    })

    it('should decrement like count when unliking', () => {
      const postWithLike = { ...mockPost, likeCount: 1 }
      usePostStore.getState().setPosts([postWithLike])

      usePostStore.getState().toggleLike('1', false)

      const state = usePostStore.getState()
      expect(state.posts[0].likeCount).toBe(0)
    })

    it('should update current post like count', () => {
      usePostStore.getState().setPosts([mockPost])
      usePostStore.getState().setCurrentPost(mockPost)

      usePostStore.getState().toggleLike('1', true)

      const state = usePostStore.getState()
      expect(state.currentPost?.likeCount).toBe(1)
    })

    it('should not affect other posts', () => {
      const post1 = { ...mockPost, id: '1' }
      const post2 = { ...mockPost, id: '2' }
      usePostStore.getState().setPosts([post1, post2])

      usePostStore.getState().toggleLike('1', true)

      const state = usePostStore.getState()
      expect(state.posts[0].likeCount).toBe(1)
      expect(state.posts[1].likeCount).toBe(0)
    })
  })

  describe('Bookmark Functionality', () => {
    it('should toggle bookmark', () => {
      usePostStore.getState().setPosts([mockPost])

      usePostStore.getState().toggleBookmark('1', true)

      // Bookmark doesn't change count, just the state
      const state = usePostStore.getState()
      expect(state.posts[0]).toBeDefined()
    })
  })
})

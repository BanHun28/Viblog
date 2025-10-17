import { create } from 'zustand'
import { Post, Category, Tag } from '@/types/post'

interface PostState {
  posts: Post[]
  currentPost: Post | null
  categories: Category[]
  tags: Tag[]
  isLoading: boolean
  error: string | null
  // Pagination
  hasNext: boolean
  nextCursor: string | null
}

interface PostActions {
  setPosts: (posts: Post[]) => void
  addPost: (post: Post) => void
  updatePost: (id: string, updates: Partial<Post>) => void
  deletePost: (id: string) => void
  setCurrentPost: (post: Post | null) => void
  setCategories: (categories: Category[]) => void
  setTags: (tags: Tag[]) => void
  setLoading: (loading: boolean) => void
  setError: (error: string | null) => void
  setPagination: (hasNext: boolean, nextCursor: string | null) => void
  clearPosts: () => void
  // Like/Bookmark
  toggleLike: (postId: string, isLiked: boolean) => void
  toggleBookmark: (postId: string, isBookmarked: boolean) => void
}

type PostStore = PostState & PostActions

export const usePostStore = create<PostStore>((set) => ({
  // State
  posts: [],
  currentPost: null,
  categories: [],
  tags: [],
  isLoading: false,
  error: null,
  hasNext: false,
  nextCursor: null,

  // Actions
  setPosts: (posts) => set({ posts }),

  addPost: (post) => set((state) => ({ posts: [post, ...state.posts] })),

  updatePost: (id, updates) =>
    set((state) => ({
      posts: state.posts.map((post) => (post.id === id ? { ...post, ...updates } : post)),
      currentPost:
        state.currentPost?.id === id
          ? { ...state.currentPost, ...updates }
          : state.currentPost,
    })),

  deletePost: (id) =>
    set((state) => ({
      posts: state.posts.filter((post) => post.id !== id),
      currentPost: state.currentPost?.id === id ? null : state.currentPost,
    })),

  setCurrentPost: (post) => set({ currentPost: post }),

  setCategories: (categories) => set({ categories }),

  setTags: (tags) => set({ tags }),

  setLoading: (loading) => set({ isLoading: loading }),

  setError: (error) => set({ error }),

  setPagination: (hasNext, nextCursor) => set({ hasNext, nextCursor }),

  clearPosts: () => set({ posts: [], hasNext: false, nextCursor: null }),

  toggleLike: (postId, isLiked) =>
    set((state) => ({
      posts: state.posts.map((post) =>
        post.id === postId
          ? { ...post, likeCount: post.likeCount + (isLiked ? 1 : -1) }
          : post
      ),
      currentPost:
        state.currentPost?.id === postId
          ? { ...state.currentPost, likeCount: state.currentPost.likeCount + (isLiked ? 1 : -1) }
          : state.currentPost,
    })),

  toggleBookmark: (postId, isBookmarked) =>
    set((state) => ({
      posts: state.posts.map((post) => (post.id === postId ? { ...post } : post)),
    })),
}))

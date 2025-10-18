import { apiClient } from './client'
import {
  Post,
  CreatePostRequest,
  UpdatePostRequest,
  PostListParams,
  Category,
  Tag,
} from '@/types/post'
import { PaginatedResponse, ApiResponse } from '@/types/api'

export const postsApi = {
  // Get paginated posts list
  async getPosts(params?: PostListParams): Promise<PaginatedResponse<Post>> {
    const queryParams = new URLSearchParams()
    if (params?.page) queryParams.append('page', params.page.toString())
    if (params?.limit) queryParams.append('limit', params.limit.toString())

    return apiClient.get(`/posts?${queryParams.toString()}`)
  },

  // Get post by ID
  async getPostById(id: string): Promise<Post> {
    const response = await apiClient.get<ApiResponse<Post>>(`/posts/${id}`)
    return response.data
  },

  // Create new post (Admin only)
  async createPost(data: CreatePostRequest): Promise<Post> {
    const response = await apiClient.post<ApiResponse<Post>>('/posts', data)
    return response.data
  },

  // Update post (Admin only)
  async updatePost(id: string, data: UpdatePostRequest): Promise<Post> {
    const response = await apiClient.put<ApiResponse<Post>>(`/posts/${id}`, data)
    return response.data
  },

  // Delete post (Admin only)
  async deletePost(id: string): Promise<void> {
    await apiClient.delete(`/posts/${id}`)
  },

  // Like/Unlike post
  async likePost(id: string): Promise<void> {
    await apiClient.post(`/posts/${id}/like`)
  },

  async unlikePost(id: string): Promise<void> {
    await apiClient.delete(`/posts/${id}/like`)
  },

  // Bookmark/Unbookmark post
  async bookmarkPost(id: string): Promise<void> {
    await apiClient.post(`/posts/${id}/bookmark`)
  },

  async unbookmarkPost(id: string): Promise<void> {
    await apiClient.delete(`/posts/${id}/bookmark`)
  },

  // Increment view count
  async incrementView(id: string): Promise<void> {
    await apiClient.post(`/posts/${id}/view`)
  },

  // Search posts
  async searchPosts(query: string): Promise<Post[]> {
    const response = await apiClient.get<ApiResponse<Post[]>>(`/posts/search?q=${encodeURIComponent(query)}`)
    return response.data
  },

  // Get categories
  async getCategories(): Promise<Category[]> {
    const response = await apiClient.get<ApiResponse<Category[]>>('/categories')
    return response.data
  },

  // Get posts by category
  async getPostsByCategory(slug: string, params?: PostListParams): Promise<PaginatedResponse<Post>> {
    const queryParams = new URLSearchParams()
    if (params?.page) queryParams.append('page', params.page.toString())
    if (params?.limit) queryParams.append('limit', params.limit.toString())

    return apiClient.get(`/categories/${slug}/posts?${queryParams.toString()}`)
  },

  // Get tags
  async getTags(): Promise<Tag[]> {
    const response = await apiClient.get<ApiResponse<Tag[]>>('/tags')
    return response.data
  },

  // Get posts by tag
  async getPostsByTag(slug: string, params?: PostListParams): Promise<PaginatedResponse<Post>> {
    const queryParams = new URLSearchParams()
    if (params?.page) queryParams.append('page', params.page.toString())
    if (params?.limit) queryParams.append('limit', params.limit.toString())

    return apiClient.get(`/tags/${slug}/posts?${queryParams.toString()}`)
  },
}

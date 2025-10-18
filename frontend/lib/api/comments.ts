import { apiClient } from './client'
import { Comment } from '@/types/comment'
import { ApiResponse } from '@/types/api'

export const commentsApi = {
  // Get comments by post ID
  async getCommentsByPostId(postId: number | string): Promise<Comment[]> {
    const response = await apiClient.get<ApiResponse<Comment[]>>(`/comments/post/${postId}`)
    return response.data
  },

  // Get replies for a comment
  async getReplies(commentId: number | string): Promise<Comment[]> {
    const response = await apiClient.get<ApiResponse<Comment[]>>(`/comments/${commentId}/replies`)
    return response.data
  },

  // Create a comment
  async createComment(postId: number | string, data: { content: string; authorName?: string; authorEmail?: string }): Promise<Comment> {
    const response = await apiClient.post<ApiResponse<Comment>>(`/comments/post/${postId}`, data)
    return response.data
  },

  // Create a reply
  async createReply(commentId: number | string, data: { content: string; authorName?: string; authorEmail?: string }): Promise<Comment> {
    const response = await apiClient.post<ApiResponse<Comment>>(`/comments/${commentId}/replies`, data)
    return response.data
  },

  // Update a comment
  async updateComment(id: number | string, data: { content: string }): Promise<Comment> {
    const response = await apiClient.put<ApiResponse<Comment>>(`/comments/${id}`, data)
    return response.data
  },

  // Delete a comment
  async deleteComment(id: number | string): Promise<void> {
    await apiClient.delete(`/comments/${id}`)
  },

  // Like a comment
  async likeComment(id: number | string): Promise<void> {
    await apiClient.post(`/comments/${id}/like`)
  },

  // Unlike a comment
  async unlikeComment(id: number | string): Promise<void> {
    await apiClient.delete(`/comments/${id}/like`)
  },
}
